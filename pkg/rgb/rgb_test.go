package rgb

import "testing"

func TestFrameBrightness(t *testing.T) {
	s := New(2)
	s.LED(0).R = 100
	s.LED(1).G = 200
	s.SetBrightness(128)
	frame := s.Frame()
	if frame[0].R != 50 {
		t.Errorf("expected red 50, got %d", frame[0].R)
	}
	if frame[1].G != 100 {
		t.Errorf("expected green 100, got %d", frame[1].G)
	}
}

func TestLEDAccess(t *testing.T) {
	s := New(1)
	led := s.LED(0)
	if led == nil {
		t.Fatalf("expected LED, got nil")
	}
	led.B = 255
	frame := s.Frame()
	if frame[0].B != 255 {
		t.Errorf("expected blue 255, got %d", frame[0].B)
	}
}
