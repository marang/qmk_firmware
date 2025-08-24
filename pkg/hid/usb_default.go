//go:build tinygo

package hid

import "machine"

// DefaultUSB returns the machine USB device.
func DefaultUSB() USBDevice {
	return machine.USB
}
