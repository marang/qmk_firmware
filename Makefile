TINYGO ?= tinygo
TARGET ?= ./board/keychron_q1_pro.json
OUT_DIR ?= build
OUT_BIN ?= $(OUT_DIR)/firmware.elf
TINYGOROOT := $(shell $(TINYGO) env TINYGOROOT)
LOCAL_TINYGOROOT := $(OUT_DIR)/tinygoroot
ABS_TINYGOROOT := $(abspath $(LOCAL_TINYGOROOT))

.PHONY: build clean flash

build:
	mkdir -p $(OUT_DIR)
	cp -r $(TINYGOROOT) $(LOCAL_TINYGOROOT)
	cp machine/keychron_q1_pro.go $(LOCAL_TINYGOROOT)/src/machine/keychron_q1_pro.go
	TINYGOROOT=$(ABS_TINYGOROOT) $(TINYGO) build -target $(TARGET) -o $(OUT_BIN) ./cmd/firmware

clean:
	rm -rf $(OUT_DIR)

flash: build
	go run ./cmd/flash -bin $(OUT_BIN)
