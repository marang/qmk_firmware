package keymap

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/qmk/qmk_firmware/pkg/hid"
)

func TestParseJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "keymap.json")
	data := `{
        "layers": [{"name": "base", "keys": [["KC_A", "KC_B"]]}],
        "macros": {"test": ["KC_A", "KC_B"]}
    }`
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	km, err := LoadJSON(path)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if len(km.Layers) != 1 || km.Layers[0].Keys[0][0] != hid.KeyA {
		t.Fatalf("unexpected keymap: %+v", km)
	}
	if seq := km.Macros["test"]; len(seq) != 2 || seq[1] != hid.KeyB {
		t.Fatalf("unexpected macro: %+v", seq)
	}
}

func TestMacroRecorder(t *testing.T) {
	mr := NewMacroRecorder()
	mr.Start("hello")
	mr.Record(hid.KeyA)
	mr.Record(hid.KeyB)
	mr.Stop()
	got := mr.Play("hello")
	want := []hid.Keycode{hid.KeyA, hid.KeyB}
	for i, k := range want {
		if got[i] != k {
			t.Fatalf("macro mismatch: got %v want %v", got, want)
		}
	}
}

func TestParseAdditionalKeycodes(t *testing.T) {
	data := `{
               "layers": [{"name": "base", "keys": [["KC_1", "KC_RALT"], ["KC_F5", "KC_Z"]]}],
               "macros": {"combo": ["KC_LCTL", "KC_0", "KC_RSFT"]}
       }`
	km, err := ParseJSON(strings.NewReader(data))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if kc := km.Layers[0].Keys[0][0]; kc != hid.Keycode(0x1e) {
		t.Fatalf("KC_1 -> %v", kc)
	}
	if kc := km.Layers[0].Keys[0][1]; kc != hid.Keycode(0xe6) {
		t.Fatalf("KC_RALT -> %v", kc)
	}
	if kc := km.Layers[0].Keys[1][0]; kc != hid.Keycode(0x3e) {
		t.Fatalf("KC_F5 -> %v", kc)
	}
	if kc := km.Layers[0].Keys[1][1]; kc != hid.Keycode(0x1d) {
		t.Fatalf("KC_Z -> %v", kc)
	}
	macro := km.Macros["combo"]
	want := []hid.Keycode{0xe0, 0x27, 0xe5}
	for i, k := range want {
		if macro[i] != k {
			t.Fatalf("macro mismatch: got %v want %v", macro, want)
		}
	}
}
