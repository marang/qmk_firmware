TINYGO ?= tinygo
TARGET ?= ./board/keychron_q1_pro.json
OUT_DIR ?= build
OUT_BIN ?= $(OUT_DIR)/firmware.elf
TINYGOROOT := $(shell $(TINYGO) env TINYGOROOT)

.PHONY: build clean flash

build:
	mkdir -p $(OUT_DIR)
	cp machine/keychron_q1_pro.go $(TINYGOROOT)/src/machine/keychron_q1_pro.go
	$(TINYGO) build -target $(TARGET) -o $(OUT_BIN) ./cmd/firmware

clean:
	rm -rf $(OUT_DIR)

flash:
	mkdir -p $(OUT_DIR)
	$(TINYGO) build -target $(TARGET) -o $(OUT_BIN) ./cmd/firmware
	go run ./cmd/flash -bin $(OUT_BIN)
