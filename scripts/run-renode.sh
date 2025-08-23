#!/bin/bash
set -euo pipefail

FIRMWARE=${FIRMWARE:-build/firmware.elf}
RENODE_LOG=${RENODE_LOG:-renode.log}

make build

renode --disable-xwt \
  -e 'mach create "stm32l432-keyboard"' \
  -e 'machine LoadPlatformDescription @renode/stm32l432_keyboard.repl' \
  -e "sysbus LoadELF @$FIRMWARE" \
  -e 'connector Connect sysbus.usb 1234' \
  -e 'start' \
  2>&1 | tee "$RENODE_LOG"
