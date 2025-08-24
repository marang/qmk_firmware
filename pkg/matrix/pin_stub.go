//go:build !tinygo

package matrix

type PinMode uint8

type PinConfig struct{ Mode PinMode }

const (
	PinInput PinMode = iota
	PinOutput
)
