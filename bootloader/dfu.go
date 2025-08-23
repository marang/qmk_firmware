package bootloader

import (
	"debug/elf"
	"fmt"
	"os"
	"os/exec"
)

const (
	flashStart = 0x08000000
	flashSize  = 256 * 1024
	flashEnd   = flashStart + flashSize
)

func isInFlash(addr, size uint64) bool {
	if addr < flashStart {
		return false
	}
	return addr+size <= flashEnd
}

func verifySection(addr, size uint64) error {
	if !isInFlash(addr, size) {
		return fmt.Errorf("section 0x%X size 0x%X outside flash", addr, size)
	}
	return nil
}

// VerifyLayout ensures all loadable sections fit within flash memory and the entry point is valid.
func VerifyLayout(path string) error {
	f, err := elf.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if !isInFlash(uint64(f.Entry), 1) {
		return fmt.Errorf("entry point 0x%X outside flash", f.Entry)
	}
	for _, s := range f.Sections {
		if s.Flags&elf.SHF_ALLOC != 0 {
			if err := verifySection(s.Addr, s.Size); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flash verifies the firmware layout and uploads it using dfu-util.
func Flash(path string) error {
	if err := VerifyLayout(path); err != nil {
		return err
	}
	cmd := exec.Command("dfu-util", "-a", "0", "-D", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
