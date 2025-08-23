package power

import "testing"

func TestStopWakeGPIO(t *testing.T) {
	Stop()
	if State() != ModeStop {
		t.Fatalf("expected stop mode, got %v", State())
	}
	WakeGPIO()
	if State() != ModeActive {
		t.Fatalf("expected active after GPIO wake, got %v", State())
	}
}

func TestStandbyWakeUSB(t *testing.T) {
	Standby()
	if State() != ModeStandby {
		t.Fatalf("expected standby mode, got %v", State())
	}
	WakeUSB()
	if State() != ModeActive {
		t.Fatalf("expected active after USB wake, got %v", State())
	}
}
