package power

// Mode represents MCU power mode.
type Mode int

const (
	ModeActive Mode = iota
	ModeStop
	ModeStandby
)

var state Mode = ModeActive

// Stop puts the MCU into stop mode.
func Stop() { state = ModeStop }

// Standby puts the MCU into standby mode.
func Standby() { state = ModeStandby }

// WakeGPIO wakes the MCU due to a GPIO interrupt.
func WakeGPIO() { state = ModeActive }

// WakeUSB wakes the MCU due to USB activity.
func WakeUSB() { state = ModeActive }

// State returns the current power state.
func State() Mode { return state }
