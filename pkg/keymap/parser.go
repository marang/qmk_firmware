package keymap

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/qmk/qmk_firmware/pkg/hid"
)

// keycodeLookup maps string names to HID keycodes.
var keycodeLookup = map[string]hid.Keycode{
	"KC_A": hid.KeyA,
	"KC_B": hid.KeyB,
}

type jsonLayer struct {
	Name string     `json:"name"`
	Keys [][]string `json:"keys"`
}

type jsonKeymap struct {
	Layers []jsonLayer         `json:"layers"`
	Macros map[string][]string `json:"macros"`
}

// ParseJSON reads a keymap from the given reader.
func ParseJSON(r io.Reader) (*Keymap, error) {
	var jk jsonKeymap
	if err := json.NewDecoder(r).Decode(&jk); err != nil {
		return nil, err
	}
	km := &Keymap{Macros: make(map[string][]hid.Keycode)}
	for _, l := range jk.Layers {
		layer := Layer{Name: l.Name}
		for _, row := range l.Keys {
			var keyRow []hid.Keycode
			for _, key := range row {
				kc, ok := keycodeLookup[key]
				if !ok {
					return nil, fmt.Errorf("unknown keycode %q", key)
				}
				keyRow = append(keyRow, kc)
			}
			layer.Keys = append(layer.Keys, keyRow)
		}
		km.Layers = append(km.Layers, layer)
	}
	for name, seq := range jk.Macros {
		for _, key := range seq {
			kc, ok := keycodeLookup[key]
			if !ok {
				return nil, fmt.Errorf("unknown keycode %q", key)
			}
			km.Macros[name] = append(km.Macros[name], kc)
		}
	}
	return km, nil
}

// LoadJSON opens a file and parses the keymap.
func LoadJSON(path string) (*Keymap, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseJSON(f)
}
