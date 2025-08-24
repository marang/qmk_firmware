//go:build keychron_q1_pro

package machine

const (
	// SPI1 pins.
	SPI0_SCK_PIN = PA5
	SPI0_SDO_PIN = PA6
	SPI0_SDI_PIN = PA7

	// I2C1 pins.
	I2C0_SCL_PIN = PB8
	I2C0_SDA_PIN = PB9

	// UART1 pins.
	UART_TX_PIN = PA9
	UART_RX_PIN = PA10

	// USB FS pins.
	USBCDC_DM_PIN = PA11
	USBCDC_DP_PIN = PA12
)

type Serialer interface {
	Configure()
	Write([]byte) (int, error)
	WriteByte(byte) error
	ReadByte() (byte, error)
	Buffered() int
}

type USBCDCType struct{}

var (
	USB    USBCDCType
	USBCDC Serialer = USB
)

func (USBCDCType) Configure() {}

func (USBCDCType) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func (USBCDCType) WriteByte(b byte) error { return nil }

func (USBCDCType) ReadByte() (byte, error) { return 0, nil }

func (USBCDCType) Buffered() int { return 0 }

var (
	SPI0 SPI
	I2C0 I2C
)
