package keymap // import "github.com/qmk/qmk_firmware/pkg/keymap"

type Keymap struct{ ... }
    func LoadJSON(path string) (*Keymap, error)
    func ParseJSON(r io.Reader) (*Keymap, error)
type Layer struct{ ... }
type MacroRecorder struct{ ... }
    func NewMacroRecorder() *MacroRecorder
