package bootloader

import "testing"

func TestVerifySection(t *testing.T) {
	if err := verifySection(flashStart, 16); err != nil {
		t.Fatalf("expected ok, got %v", err)
	}
	if err := verifySection(flashEnd-16, 16); err != nil {
		t.Fatalf("expected ok at end, got %v", err)
	}
	if err := verifySection(flashEnd-15, 16); err == nil {
		t.Fatalf("expected error for section beyond flash")
	}
}

func TestIsInFlash(t *testing.T) {
	if !isInFlash(flashStart, 1) {
		t.Fatal("flash start should be valid")
	}
	if isInFlash(flashStart-1, 1) {
		t.Fatal("address before flash should be invalid")
	}
}
