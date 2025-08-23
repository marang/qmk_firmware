QEMU_SCRIPT ?= util/qemu_stm32f4.sh

qemu: build
	$(QEMU_SCRIPT) $(BUILD_DIR)/$(TARGET).elf
