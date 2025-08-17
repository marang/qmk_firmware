# Set proper platform and bootloader for STM32
MCU = STM32L432
BOOTLOADER = stm32-dfu

# Enter lower-power sleep mode when on the ChibiOS idle thread
OPT_DEFS += -DCORTEX_ENABLE_WFI_IDLE=TRUE
OPT_DEFS += -DNO_USB_STARTUP_CHECK -DENABLE_FACTORY_TEST

include keyboards/keychron/bluetooth/bluetooth.mk
include keyboards/keychron/common/common.mk