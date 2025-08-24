//go:build debug
// +build debug

package debug

import (
	"fmt"
	machine "github.com/qmk/qmk_firmware/machine"
)

// Printf writes formatted debug output to the USB CDC interface.
func Printf(format string, a ...interface{}) {
	machine.USB.Write([]byte(fmt.Sprintf(format, a...)))
}
