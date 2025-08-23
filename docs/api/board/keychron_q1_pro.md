package keychron_q1_pro // import "github.com/qmk/qmk_firmware/board/keychron_q1_pro"

const (
    MatrixCol0   machine.Pin
    MatrixCol1   machine.Pin
    MatrixRow0   machine.Pin
    MatrixRow1   machine.Pin
    MatrixRow2   machine.Pin
    SPI_SCK      machine.Pin
    SPI_MISO     machine.Pin
    SPI_MOSI     machine.Pin
    I2C_SCL      machine.Pin
    I2C_SDA      machine.Pin
    USB_DM       machine.Pin
    USB_DP       machine.Pin
    USB_PULLUP   machine.Pin
)

// Init sets matrix column and row modes, assigns alternate functions for
// SPI1, I2C1, and USB pins, enables peripheral clocks, and activates the USB
// pull-up resistor.
func Init()
