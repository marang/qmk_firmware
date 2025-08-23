# Hey Emacs, this is a -*- makefile -*-

RENODE ?= renode
RENODE_RESC ?=
RENODE_ARGS ?=

# If no .resc path was provided, try to use one from the keyboard directory
ifeq ($(RENODE_RESC),)
  ifneq ("$(wildcard $(KEYBOARD_PATH)/renode.resc)","")
    RENODE_RESC := $(KEYBOARD_PATH)/renode.resc
  endif
endif

# Launch Renode with the built ELF and the provided script
renode: $(BUILD_DIR)/$(TARGET).elf
	$(RENODE) $(RENODE_ARGS) $(RENODE_RESC) $(BUILD_DIR)/$(TARGET).elf
