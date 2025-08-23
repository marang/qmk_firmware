package keymap

import "github.com/qmk/qmk_firmware/pkg/hid"

// Layer represents a single keymap layer with a name and a grid of keycodes.
type Layer struct {
	Name string
	Keys [][]hid.Keycode
}

// Keymap contains all layers and any defined macros.
type Keymap struct {
	Layers []Layer
	Macros map[string][]hid.Keycode
}

// MacroRecorder provides a simple API for recording and playing back macros.
type MacroRecorder struct {
	Macros    map[string][]hid.Keycode
	recording []hid.Keycode
	current   string
}

// NewMacroRecorder returns an initialised MacroRecorder.
func NewMacroRecorder() *MacroRecorder {
	return &MacroRecorder{Macros: make(map[string][]hid.Keycode)}
}

// Start begins recording a macro under the given name.
func (m *MacroRecorder) Start(name string) {
	m.current = name
	m.recording = nil
}

// Record appends a keycode to the active recording.
func (m *MacroRecorder) Record(k hid.Keycode) {
	if m.current != "" {
		m.recording = append(m.recording, k)
	}
}

// Stop finishes the current recording and stores it by name.
func (m *MacroRecorder) Stop() {
	if m.current != "" {
		m.Macros[m.current] = append([]hid.Keycode(nil), m.recording...)
		m.current = ""
		m.recording = nil
	}
}

// Play returns a copy of the macro sequence with the given name.
func (m *MacroRecorder) Play(name string) []hid.Keycode {
	return append([]hid.Keycode(nil), m.Macros[name]...)
}
