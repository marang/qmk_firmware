//go:build tinygo

package main

import "github.com/qmk/qmk_firmware/board/keychron_q1_pro"

func main() {
	keychron_q1_pro.Init()
	for {
	}
}
