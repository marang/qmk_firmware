package machine // import "github.com/qmk/qmk_firmware/machine"

const CPUFrequency = 80_000_000
type I2C struct{}
    var I2C0 I2C
type I2CConfig struct{}
type Pin uint8
    const PA0 Pin = iota ...
type PinConfig struct{ ... }
type PinMode uint8
    const PinInput PinMode = iota ...
type SPI struct{}
    var SPI0 SPI
type SPIConfig struct{}
type USBCDC struct{}
    var USB USBCDC
