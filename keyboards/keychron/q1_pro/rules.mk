# Enter lower-power sleep mode when on the ChibiOS idle thread
OPT_DEFS += -DCORTEX_ENABLE_WFI_IDLE=TRUE
OPT_DEFS += -DNO_USB_STARTUP_CHECK

BLUETOOTH_ENABLE = yes
BLUETOOTH_DRIVER = custom
DEBOUNCE_TYPE = custom

include keyboards/keychron/common/common.mk

VIA_ENABLE = yes
