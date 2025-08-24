//go:build tinygo

package matrix

import "machine"

type PinConfig = machine.PinConfig

type PinMode = machine.PinMode

const (
	PinInput  PinMode = machine.PinInput
	PinOutput PinMode = machine.PinOutput
)
