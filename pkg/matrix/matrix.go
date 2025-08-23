package matrix

import (
	"fmt"

	"github.com/qmk/qmk_firmware/machine"
)

// Pin abstracts a GPIO pin for scanning.
type Pin interface {
	Configure(machine.PinConfig)
	Set(bool)
	Get() bool
}

type shiftRegister struct {
	data  Pin
	clock Pin
	latch Pin
}

func (sr *shiftRegister) write(value uint8) {
	for i := 0; i < 8; i++ {
		bit := (value >> uint(7-i)) & 1
		sr.data.Set(bit == 1)
		sr.clock.Set(false)
		sr.clock.Set(true)
	}
	sr.latch.Set(false)
	sr.latch.Set(true)
}

// Matrix scans a keyboard matrix using a 74HC595 shift register.
type Matrix struct {
	sr   shiftRegister
	rows []Pin
	cols int
}

// New creates a Matrix and configures the GPIO pins.
func New(data, clock, latch Pin, rows []Pin, cols int) *Matrix {
	data.Configure(machine.PinConfig{Mode: machine.PinOutput})
	clock.Configure(machine.PinConfig{Mode: machine.PinOutput})
	latch.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for _, r := range rows {
		r.Configure(machine.PinConfig{Mode: machine.PinInput})
	}
	return &Matrix{
		sr:   shiftRegister{data: data, clock: clock, latch: latch},
		rows: rows,
		cols: cols,
	}
}

// Scan advances through each column, reading row pins and logging results.
func (m *Matrix) Scan() [][]bool {
	result := make([][]bool, len(m.rows))
	for i := range result {
		result[i] = make([]bool, m.cols)
	}
	for col := 0; col < m.cols; col++ {
		m.sr.write(1 << uint(col))
		for rowIdx, row := range m.rows {
			pressed := !row.Get()
			result[rowIdx][col] = pressed
			fmt.Printf("row %d col %d: %v\n", rowIdx, col, pressed)
		}
	}
	return result
}
