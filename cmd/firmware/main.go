//go:build tinygo

package main

import (
	"time"
	"github.com/qmk/qmk_firmware/board/keychron_q1_pro"
	"github.com/qmk/qmk_firmware/pkg/debug"
)

func main() {
	// Initialize the board
	keychron_q1_pro.Init()
	
	// Print startup message
	debug.Printf("Keychron Q1 Pro firmware initialized\n")
	
	// Add a small delay to allow Renode to process initialization
	time.Sleep(time.Millisecond * 100)
	
	// Main loop - this is what Renode should be able to observe
	for i := 0; ; i++ {
		if i%10 == 0 {
			debug.Printf("System running - iteration %d\n", i)
		}
		time.Sleep(time.Second * 1)
	}
}
