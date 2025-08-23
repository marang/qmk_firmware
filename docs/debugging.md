# Debug Logging

The `debug` package provides a simple `Printf` function that writes formatted
output to the board's UART or USB CDC interface. This allows firmware to emit
useful diagnostics while running on hardware.

## Usage

Import the package and replace calls to `fmt.Printf` with `debug.Printf`:

```go
import "github.com/qmk/qmk_firmware/pkg/debug"

debug.Printf("hello %s\n", "world")
```

By default logging is disabled so production builds incur no overhead. To enable
logging, build the firmware with the `debug` build tag:

```
tinygo build -tags debug
```

When enabled, logs are sent over the board's USB CDC or UART port and can be
viewed with any serial terminal.
