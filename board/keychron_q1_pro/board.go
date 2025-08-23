package keychron_q1_pro

import (
	"device/stm32"
	"github.com/qmk/qmk_firmware/machine"
)

const (
	// Matrix column pins
	MatrixCol0 = machine.PA0
	MatrixCol1 = machine.PA1

	// Matrix row pins
	MatrixRow0 = machine.PA2
	MatrixRow1 = machine.PB0
	MatrixRow2 = machine.PB1

	// SPI1 pins
	SPI_SCK  = machine.Pin(5) // PA5
	SPI_MISO = machine.Pin(6) // PA6
	SPI_MOSI = machine.Pin(7) // PA7

	// I2C1 pins
	I2C_SCL = machine.Pin(24) // PB8
	I2C_SDA = machine.Pin(25) // PB9

	// USB pins and pull-up control
	USB_DM     = machine.Pin(11) // PA11
	USB_DP     = machine.Pin(12) // PA12
	USB_PULLUP = machine.Pin(13) // PA13
)

// setAltFunc configures a pin in alternate function mode with the provided AF number.
// If openDrain is true, the pin is configured as open-drain.
func setAltFunc(p machine.Pin, af uint32, openDrain bool) {
	// Ensure GPIO clock enabled
	p.Configure(machine.PinConfig{Mode: machine.PinInput})

	port := gpioPort(p)
	pin := gpioPin(p)

	// Set mode to alternate function (10)
	shift := uint8(pin * 2)
	port.MODER.ReplaceBits(0x2, 0x3, shift)

	if openDrain {
		port.OTYPER.SetBits(uint32(1) << pin)
	} else {
		port.OTYPER.ClearBits(uint32(1) << pin)
	}

	// Set alternate function number
	if pin < 8 {
		port.AFRL.ReplaceBits(af, 0xF, uint8(pin*4))
	} else {
		port.AFRH.ReplaceBits(af, 0xF, uint8((pin-8)*4))
	}
}

func gpioPort(p machine.Pin) *stm32.GPIO_Type {
	switch uint8(p) / 16 {
	case 0:
		return stm32.GPIOA
	case 1:
		return stm32.GPIOB
	default:
		panic("invalid port")
	}
}

func gpioPin(p machine.Pin) uint32 {
	return uint32(uint8(p) % 16)
}

// Init configures the peripherals used on the Keychron Q1 Pro.
func Init() {
	// Matrix setup
	MatrixCol0.Configure(machine.PinConfig{Mode: machine.PinOutputOpenDrain})
	MatrixCol1.Configure(machine.PinConfig{Mode: machine.PinOutputOpenDrain})
	MatrixRow0.Configure(machine.PinConfig{Mode: machine.PinInput})
	MatrixRow1.Configure(machine.PinConfig{Mode: machine.PinInput})
	MatrixRow2.Configure(machine.PinConfig{Mode: machine.PinInput})

	// SPI pins (AF5)
	setAltFunc(SPI_SCK, 5, false)
	setAltFunc(SPI_MISO, 5, false)
	setAltFunc(SPI_MOSI, 5, false)

	// I2C pins (AF4)
	setAltFunc(I2C_SCL, 4, true)
	setAltFunc(I2C_SDA, 4, true)

	// USB pins (AF10)
	setAltFunc(USB_DM, 10, false)
	setAltFunc(USB_DP, 10, false)
	USB_PULLUP.Configure(machine.PinConfig{Mode: machine.PinOutput})
	USB_PULLUP.Set(true)

	// Configure peripherals
	machine.SPI0.Configure(machine.SPIConfig{})
	machine.I2C0.Configure(machine.I2CConfig{})
	machine.USB.Configure()
}
