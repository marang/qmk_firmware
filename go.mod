module github.com/qmk/qmk_firmware

go 1.24.3

require (
    device/stm32 v0.0.0
    runtime/volatile v0.0.0
)

replace device/stm32 => ./device/stm32
replace runtime/volatile => ./runtime/volatile

