package keymap

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/qmk/qmk_firmware/pkg/hid"
)

// keycodeLookup maps string names to HID keycodes.
var keycodeLookup = func() map[string]hid.Keycode {
	m := map[string]hid.Keycode{
		// Punctuation and control keys
		"KC_ENT":  0x28,
		"KC_ESC":  0x29,
		"KC_BSPC": 0x2a,
		"KC_TAB":  0x2b,
		"KC_SPC":  0x2c,
		"KC_MINS": 0x2d,
		"KC_EQL":  0x2e,
		"KC_LBRC": 0x2f,
		"KC_RBRC": 0x30,
		"KC_BSLS": 0x31,
		"KC_NUHS": 0x32,
		"KC_SCLN": 0x33,
		"KC_QUOT": 0x34,
		"KC_GRV":  0x35,
		"KC_COMM": 0x36,
		"KC_DOT":  0x37,
		"KC_SLSH": 0x38,
		"KC_CAPS": 0x39,
	}
	// Letters
	for i := 0; i < 26; i++ {
		m[fmt.Sprintf("KC_%c", 'A'+i)] = hid.Keycode(0x04 + i)
	}
	// Numbers
	for i, name := range []string{"KC_1", "KC_2", "KC_3", "KC_4", "KC_5", "KC_6", "KC_7", "KC_8", "KC_9", "KC_0"} {
		m[name] = hid.Keycode(0x1e + i)
	}
	// Function keys
	for i := 0; i < 24; i++ {
		m[fmt.Sprintf("KC_F%d", i+1)] = hid.Keycode(0x3a + i)
	}
	// Navigation keys
	navNames := []string{"KC_PSCR", "KC_SLCK", "KC_PAUS", "KC_INS", "KC_HOME", "KC_PGUP", "KC_DEL", "KC_END", "KC_PGDN", "KC_RGHT", "KC_LEFT", "KC_DOWN", "KC_UP"}
	navCodes := []hid.Keycode{0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52}
	for i, name := range navNames {
		m[name] = navCodes[i]
	}
	// Keypad keys
	kpNames := []string{"KC_NLCK", "KC_PSLS", "KC_PAST", "KC_PMNS", "KC_PPLS", "KC_PENT", "KC_P1", "KC_P2", "KC_P3", "KC_P4", "KC_P5", "KC_P6", "KC_P7", "KC_P8", "KC_P9", "KC_P0", "KC_PDOT", "KC_NUBS", "KC_APP", "KC_POWER", "KC_PEQL"}
	kpCodes := []hid.Keycode{0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67}
	for i, name := range kpNames {
		m[name] = kpCodes[i]
	}
	// Editing and international keys
	editNames := []string{
		"KC_EXEC", "KC_HELP", "KC_MENU", "KC_SELECT", "KC_STOP", "KC_AGAIN", "KC_UNDO", "KC_CUT", "KC_COPY", "KC_PASTE", "KC_FIND", "KC_MUTE", "KC_VOLU", "KC_VOLD",
		"KC_LOCKING_CAPS", "KC_LOCKING_NUM", "KC_LOCKING_SCROLL", "KC_KP_COMMA", "KC_KP_EQUAL_AS400",
		"KC_INT1", "KC_INT2", "KC_INT3", "KC_INT4", "KC_INT5", "KC_INT6", "KC_INT7", "KC_INT8", "KC_INT9",
		"KC_LANG1", "KC_LANG2", "KC_LANG3", "KC_LANG4", "KC_LANG5", "KC_LANG6", "KC_LANG7", "KC_LANG8", "KC_LANG9",
		"KC_ALT_ERASE", "KC_SYSREQ", "KC_CANCEL", "KC_CLEAR", "KC_PRIOR", "KC_RETURN", "KC_SEPARATOR", "KC_OUT", "KC_OPER", "KC_CLEAR_AGAIN", "KC_CRSEL", "KC_EXSEL",
	}
	editCodes := []hid.Keycode{0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f, 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86,
		0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
		0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98,
		0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f, 0xa0, 0xa1, 0xa2, 0xa3, 0xa4}
	for i, name := range editNames {
		m[name] = editCodes[i]
	}
	// Modifiers
	modNames := []string{"KC_LCTL", "KC_LSFT", "KC_LALT", "KC_LGUI", "KC_RCTL", "KC_RSFT", "KC_RALT", "KC_RGUI"}
	for i, name := range modNames {
		m[name] = hid.Keycode(0xe0 + i)
	}
	return m
}()

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
