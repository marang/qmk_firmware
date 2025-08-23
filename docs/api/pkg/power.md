package power // import "github.com/qmk/qmk_firmware/pkg/power"

func Standby()
func Stop()
func WakeGPIO()
func WakeUSB()
type Mode int
    const ModeActive Mode = iota ...
    func State() Mode
