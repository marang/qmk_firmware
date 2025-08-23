# Hey Emacs, this is a -*- makefile -*-

RENODE ?= renode
RENODE_RESC ?=
RENODE_ARGS ?=

# Determine the .resc script to use
ifeq ($(RENODE_RESC),)
  ifneq ("$(wildcard $(KEYBOARD_PATH)/renode.resc)","")
    RENODE_RESC := $(KEYBOARD_PATH)/renode.resc
  else
    RENODE_RESC := $(ROOT_DIR)/util/renode/$(subst /,_,$(KEYBOARD)).resc
  endif
endif

# Launch Renode with the built ELF and the provided script
.PHONY: renode
renode: $(BUILD_DIR)/$(TARGET).elf
	$(RENODE) $(RENODE_ARGS) "$(RENODE_RESC)" "$(BUILD_DIR)/$(TARGET).elf"
