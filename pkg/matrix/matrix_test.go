package matrix

import (
	"testing"

	"github.com/qmk/qmk_firmware/machine"
	"github.com/qmk/qmk_firmware/pkg/rgb"
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

	strip := rgb.New(1)
	strip.SetBrightness(128)
	led := strip.LED(0)
	led.R = 255
	frame := strip.Frame()
	if frame[0].R != 128 {
		t.Fatalf("expected red 128, got %d", frame[0].R)
	}
	m.Scan()
	led.R, led.G = 0, 255
	frame = strip.Frame()
	if frame[0].G != 128 {
		t.Fatalf("expected green 128 after toggle, got %d", frame[0].G)
	}
}
