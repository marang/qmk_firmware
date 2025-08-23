package hid

import (
	"fmt"
	"github.com/qmk/qmk_firmware/machine"
)

// Keycode represents a USB HID key usage ID.
type Keycode uint8

// Common keycodes used in tests.
const (
	KeyA Keycode = 0x04
	KeyB Keycode = 0x05
)

// Report is an 8-byte keyboard report: modifier, reserved, 6 keycodes.
type Report struct {
	Modifier byte
	Keys     [6]byte
}

// Bytes returns the raw report data.
func (r Report) Bytes() []byte {
	data := make([]byte, 8)
	data[0] = r.Modifier
	copy(data[2:], r.Keys[:])
	return data
}

// USBDevice is the minimal interface needed for HID reports.
type USBDevice interface {
	Configure()
	Write([]byte) (int, error)
}

// Keyboard maps matrix events to HID reports and sends them over USB.
type Keyboard struct {
	usb    USBDevice
	keymap map[[2]int]Keycode
}

// NewKeyboard configures the USB device and returns a Keyboard.
func NewKeyboard(usb USBDevice, keymap map[[2]int]Keycode) *Keyboard {
	usb.Configure()
	fmt.Println("usb: enumerated")
	return &Keyboard{usb: usb, keymap: keymap}
}

// ProcessMatrix converts pressed keys to a HID report and logs events.
func (k *Keyboard) ProcessMatrix(matrix [][]bool) Report {
	var report Report
	idx := 0
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] {
				if kc, ok := k.keymap[[2]int{r, c}]; ok && idx < len(report.Keys) {
					report.Keys[idx] = byte(kc)
					fmt.Printf("key r%d c%d -> 0x%X\n", r, c, kc)
					idx++
				}
			}
		}
	}
	k.usb.Write(report.Bytes())
	return report
}

// DefaultUSB returns the machine USB device.
func DefaultUSB() USBDevice {
	return machine.USB
}
