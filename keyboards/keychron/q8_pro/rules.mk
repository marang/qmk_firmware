# Build Options
#   change yes to no to disable
#
# Enter lower-power sleep mode when on the ChibiOS idle thread
OPT_DEFS += -DCORTEX_ENABLE_WFI_IDLE=TRUE
OPT_DEFS += -DNO_USB_STARTUP_CHECK -DENABLE_FACTORY_TEST
OPT_DEFS += -DKC_BLUETOOTH_ENABLE

CUSTOM_MATRIX = lite

SRC += matrix.c

include keyboards/keychron/bluetooth/bluetooth.mk

DEFAULT_FOLDER = keychron/q8_pro/ansi_encoder
