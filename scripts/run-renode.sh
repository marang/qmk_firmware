#!/bin/bash
set -euo pipefail

FIRMWARE=${FIRMWARE:-build/firmware.elf}
RENODE_LOG=${RENODE_LOG:-renode.log}
RENODE_BIN=${RENODE:-renode}

if ! command -v "$RENODE_BIN" >/dev/null 2>&1; then
  echo "Error: renode command not found. Install renode or set RENODE to the executable path." >&2
  exit 1
fi

make build

"$RENODE_BIN" --disable-xwt \
  -e 'mach create "stm32l432-keyboard"' \
  -e 'machine LoadPlatformDescription @renode/stm32l432_keyboard.repl' \
  -e "sysbus LoadELF @$FIRMWARE" \
  -e 'connector Connect sysbus.usb 1234' \
  -e 'start' \
  2>&1 | tee "$RENODE_LOG"
