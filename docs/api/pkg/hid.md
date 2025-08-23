package hid // import "github.com/qmk/qmk_firmware/pkg/hid"

type Keyboard struct{ ... }
    func NewKeyboard(usb USBDevice, keymap map[[2]int]Keycode) *Keyboard
type Keycode uint8
    const KeyA Keycode = 0x04 ...
type Report struct{ ... }
type USBDevice interface{ ... }
    func DefaultUSB() USBDevice
