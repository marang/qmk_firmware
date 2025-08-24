module github.com/qmk/qmk_firmware

go 1.24.3

require github.com/qmk/qmk_firmware/device/stm32 v0.0.0

require github.com/qmk/qmk_firmware/runtime/volatile v0.0.0 // indirect

replace github.com/qmk/qmk_firmware/device/stm32 => ./device/stm32

replace github.com/qmk/qmk_firmware/runtime/volatile => ./runtime/volatile
