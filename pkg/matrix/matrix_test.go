package matrix

import (
	"testing"

	"github.com/qmk/qmk_firmware/machine"
)

type testPin struct {
	state bool
}

func (p *testPin) Configure(machine.PinConfig) {}
func (p *testPin) Set(b bool)                  { p.state = b }
func (p *testPin) Get() bool                   { return p.state }

func TestScan(t *testing.T) {
	rows := []Pin{&testPin{state: true}, &testPin{state: false}}
	m := New(&testPin{}, &testPin{}, &testPin{}, rows, 2)
	got := m.Scan()
	if got[0][0] || got[0][1] {
		t.Errorf("row 0 expected all false, got %v", got[0])
	}
	if !got[1][0] || !got[1][1] {
		t.Errorf("row 1 expected all true, got %v", got[1])
	}
}
