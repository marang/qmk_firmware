package main

import (
	"flag"
	"log"

	"github.com/qmk/qmk_firmware/bootloader"
)

func main() {
	bin := flag.String("bin", "", "path to firmware binary")
	flag.Parse()
	if *bin == "" {
		log.Fatal("-bin is required")
	}
	if err := bootloader.Flash(*bin); err != nil {
		log.Fatal(err)
	}
}
