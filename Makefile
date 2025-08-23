TINYGO ?= tinygo
TARGET ?= board/keychron_q1_pro
OUT_DIR ?= build
OUT_BIN ?= $(OUT_DIR)/firmware.elf

.PHONY: build clean

build:
	mkdir -p $(OUT_DIR)
	$(TINYGO) build -target $(TARGET) -o $(OUT_BIN) ./cmd/firmware

clean:
	rm -rf $(OUT_DIR)
