package machine

import "github.com/qmk/qmk_firmware/device/stm32"

// Basic support for the STM32L432 microcontroller. This file adds minimal
// definitions for clocks, GPIO, USB, SPI, and I2C peripherals so that TinyGo
// can reference the hardware resources present on the Keychron Q1 Pro.

// CPUFrequency is the main system clock frequency in hertz.
const CPUFrequency = 80_000_000

type Pin uint8

const (
	portA Pin = iota * 16
	portB
)

// GPIO pin numbers.
const (
	PA0  Pin = portA + 0
	PA1  Pin = portA + 1
	PA2  Pin = portA + 2
	PA3  Pin = portA + 3
	PA4  Pin = portA + 4
	PA5  Pin = portA + 5
	PA6  Pin = portA + 6
	PA7  Pin = portA + 7
	PA8  Pin = portA + 8
	PA9  Pin = portA + 9
	PA10 Pin = portA + 10
	PA11 Pin = portA + 11
	PA12 Pin = portA + 12
	PA13 Pin = portA + 13
	PA14 Pin = portA + 14
	PA15 Pin = portA + 15

	PB0 Pin = portB + 0
	PB1 Pin = portB + 1
	PB2 Pin = portB + 2
	PB3 Pin = portB + 3
	PB4 Pin = portB + 4
	PB5 Pin = portB + 5
	PB6 Pin = portB + 6
	PB7 Pin = portB + 7
	PB8 Pin = portB + 8
	PB9 Pin = portB + 9
)

// Default peripheral pin assignments.
const (
	// SPI1 pins.
	SPI0_SCK_PIN = PA5
	SPI0_SDO_PIN = PA6
	SPI0_SDI_PIN = PA7

	// I2C1 pins.
	I2C0_SCL_PIN = PB8
	I2C0_SDA_PIN = PB9

	// USB FS pins.
	USBCDC_DM_PIN = PA11
	USBCDC_DP_PIN = PA12
)

type PinConfig struct {
	Mode PinMode
}

type PinMode uint8

const (
	PinInput PinMode = iota
	PinOutput
	PinOutputOpenDrain
)

func (p Pin) port() uint8 { return uint8(p) / 16 }
func (p Pin) pin() uint8  { return uint8(p) % 16 }

func (p Pin) getPort() *stm32.GPIO_Type {
	switch p.port() {
	case 0:
		return stm32.GPIOA
	case 1:
		return stm32.GPIOB
	default:
		panic("invalid port")
	}
}

func (p Pin) enableClock() {
	switch p.port() {
	case 0:
		stm32.RCC.AHB2ENR.SetBits(stm32.RCC_AHB2ENR_GPIOAEN)
	case 1:
		stm32.RCC.AHB2ENR.SetBits(stm32.RCC_AHB2ENR_GPIOBEN)
	}
}

func (p Pin) Configure(cfg PinConfig) {
	p.enableClock()
	port := p.getPort()
	shift := p.pin() * 2
	port.MODER.ReplaceBits(uint32(cfg.Mode), 0x3, shift)
	switch cfg.Mode {
	case PinOutput:
		port.OTYPER.ClearBits(uint32(1) << p.pin())
	case PinOutputOpenDrain:
		port.OTYPER.SetBits(uint32(1) << p.pin())
	}
}

func (p Pin) Set(high bool) {
	port := p.getPort()
	mask := uint32(1) << p.pin()
	if high {
		port.BSRR.Set(mask)
		port.IDR.SetBits(mask)
	} else {
		port.BSRR.Set(mask << 16)
		port.IDR.ClearBits(mask)
	}
}

func (p Pin) Get() bool {
	port := p.getPort()
	mask := uint32(1) << p.pin()
	return port.IDR.HasBits(mask)
}

type SPI struct{}

var SPI0 SPI

type SPIConfig struct{}

func (SPI) Configure(cfg SPIConfig) {
	stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_SPI1EN)
	stm32.SPI1.CR1.SetBits(0x31)
}

type I2C struct{}

var I2C0 I2C

type I2CConfig struct{}

func (I2C) Configure(cfg I2CConfig) {
	stm32.RCC.APB1ENR1.SetBits(stm32.RCC_APB1ENR1_I2C1EN)
	stm32.I2C1.CR1.SetBits(1)
}

type USBCDC struct{}

var USB USBCDC

func (USBCDC) Configure() {
	stm32.RCC.APB1ENR1.SetBits(stm32.RCC_APB1ENR1_USBFSEN)
}

func (USBCDC) Write(b []byte) (int, error) {
	for _, c := range b {
		stm32.USBFS.EP0R.Set(uint32(c))
	}
	return len(b), nil
}

func init() {
	// Set up clocks to run at 80MHz.
}
