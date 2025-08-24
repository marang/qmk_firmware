TINYGO ?= tinygo
TARGET ?= ./board/keychron_q1_pro.json
OUT_DIR ?= build
OUT_BIN ?= $(OUT_DIR)/firmware.elf

.PHONY: build clean flash

build:
	mkdir -p $(OUT_DIR)
	$(TINYGO) build -target $(TARGET) -o $(OUT_BIN) ./cmd/firmware

clean:
	rm -rf $(OUT_DIR)

flash:
	mkdir -p $(OUT_DIR)
	$(TINYGO) build -target $(TARGET) -o $(OUT_BIN) ./cmd/firmware
	go run ./cmd/flash -bin $(OUT_BIN)
