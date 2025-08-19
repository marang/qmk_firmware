# Enter lower-power sleep mode when on the ChibiOS idle thread
OPT_DEFS += -DCORTEX_ENABLE_WFI_IDLE=TRUE
OPT_DEFS += -DNO_USB_STARTUP_CHECK
OPT_DEFS += -DENABLE_FACTORY_TEST

DEBOUNCE_TYPE = custom

include keyboards/keychron/bluetooth/bluetooth.mk
include keyboards/keychron/common/common.mk
SRC := $(filter-out $(COMMON_DIR)/matrix.c, $(SRC))
