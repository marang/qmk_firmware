package matrix

import (
	"time"

	"github.com/qmk/qmk_firmware/pkg/debug"
	"github.com/qmk/qmk_firmware/pkg/power"
)

// Pin abstracts a GPIO pin for scanning.
type Pin interface {
	Configure(PinConfig)
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
	data.Configure(PinConfig{Mode: PinOutput})
	clock.Configure(PinConfig{Mode: PinOutput})
	latch.Configure(PinConfig{Mode: PinOutput})
	for _, r := range rows {
		r.Configure(PinConfig{Mode: PinInput})
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
			debug.Printf("row %d col %d: %v\n", rowIdx, col, pressed)
		}
	}
	return result
}

// Loop continuously scans the matrix and enters a low power mode after the
// specified timeout with no key activity.
func (m *Matrix) Loop(timeout time.Duration, mode power.Mode) {
	last := time.Now()
	for {
		matrix := m.Scan()
		if anyPressed(matrix) {
			last = time.Now()
		}
		if time.Since(last) > timeout {
			switch mode {
			case power.ModeStandby:
				power.Standby()
			default:
				power.Stop()
			}
			return
		}
	}
}

func anyPressed(matrix [][]bool) bool {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] {
				return true
			}
		}
	}
	return false
}
