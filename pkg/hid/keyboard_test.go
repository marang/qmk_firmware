package hid

import "testing"

type fakeUSB struct {
	configured bool
	sent       [][]byte
}

func (f *fakeUSB) Configure() { f.configured = true }
func (f *fakeUSB) Write(b []byte) (int, error) {
	cp := append([]byte(nil), b...)
	f.sent = append(f.sent, cp)
	return len(b), nil
}

func TestProcessMatrix(t *testing.T) {
	usb := &fakeUSB{}
	km := map[[2]int]Keycode{{0, 0}: KeyA}
	kb := NewKeyboard(usb, km)
	if !usb.configured {
		t.Fatal("usb not configured")
	}
	matrix := [][]bool{{true}}
	rep := kb.ProcessMatrix(matrix)
	if rep.Keys[0] != byte(KeyA) {
		t.Fatalf("expected KeyA, got %v", rep.Keys[0])
	}
	if len(usb.sent) != 1 {
		t.Fatalf("expected 1 report, got %d", len(usb.sent))
	}
	if usb.sent[0][2] != byte(KeyA) {
		t.Fatalf("expected report to contain KeyA, got %v", usb.sent[0])
	}
}
