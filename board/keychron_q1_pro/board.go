package keychron_q1_pro

import "github.com/qmk/qmk_firmware/machine"

// Init configures the peripherals used on the Keychron Q1 Pro.
func Init() {
	machine.SPI0.Configure(machine.SPIConfig{})
	machine.I2C0.Configure(machine.I2CConfig{})
	machine.USB.Configure()
}
