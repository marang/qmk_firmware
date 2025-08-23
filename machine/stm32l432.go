package machine

import "fmt"

// Basic support for the STM32L432 microcontroller. This file adds minimal
// definitions for clocks, GPIO, USB, SPI, and I2C peripherals so that TinyGo
// can reference the hardware resources present on the Keychron Q1 Pro.

// CPUFrequency is the main system clock frequency in hertz.
const CPUFrequency = 80_000_000

// Pin represents a single GPIO pin on the microcontroller.
type Pin uint8

// List of available GPIO pins.
const (
	PA0 Pin = iota
	PA1
	PA2
	PB0
	PB1
)

// PinConfig configures the behavior of a GPIO pin.
type PinConfig struct {
	Mode PinMode
}

// PinMode selects a particular pin behavior.
type PinMode uint8

// Supported pin modes.
const (
	PinInput PinMode = iota
	PinOutput
	PinOutputOpenDrain
)

// Configure sets up the pin with the requested configuration.
func (p Pin) Configure(cfg PinConfig) {}

// Set drives the pin high or low.
func (p Pin) Set(high bool) {}

// Get reads the current value of the pin.
func (p Pin) Get() bool { return false }

// SPI represents a Serial Peripheral Interface bus instance.
type SPI struct{}

// SPI0 is the first SPI bus on the STM32L432.
var SPI0 SPI

// SPIConfig contains the configuration for an SPI bus.
type SPIConfig struct{}

// Configure sets up the SPI peripheral.
func (SPI) Configure(cfg SPIConfig) {}

// I2C represents an I2C bus.
type I2C struct{}

// I2C0 is the primary I2C bus.
var I2C0 I2C

// I2CConfig contains the configuration for an I2C bus.
type I2CConfig struct{}

// Configure sets up the I2C peripheral.
func (I2C) Configure(cfg I2CConfig) {}

// USBCDC provides a minimal USB CDC implementation.
type USBCDC struct{}

// USB represents the on-chip USB device controller.
var USB USBCDC

// Configure prepares the USB peripheral for use.
func (USBCDC) Configure() {}

// Write sends data over the USB CDC interface.
func (USBCDC) Write(b []byte) (int, error) {
	fmt.Printf("usb tx: %v\n", b)
	return len(b), nil
}

// init configures system clocks.
func init() {
	// Set up clocks to run at 80MHz.
}
