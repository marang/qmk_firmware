package machine

import "testing"

func TestPinSetGet(t *testing.T) {
	p := PA0
	p.Configure(PinConfig{Mode: PinOutput})
	p.Set(true)
	if !p.Get() {
		t.Fatalf("expected high state")
	}
	p.Set(false)
	if p.Get() {
		t.Fatalf("expected low state")
	}
}
