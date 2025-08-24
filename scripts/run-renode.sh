#!/bin/bash
set -euo pipefail

FIRMWARE=${FIRMWARE:-build/firmware.elf}
RENODE_LOG=${RENODE_LOG:-renode.log}
RENODE=${RENODE:-renode}

if ! command -v "$RENODE" >/dev/null 2>&1; then
  echo "Error: renode command not found. Install renode or set RENODE to the executable path." >&2
  exit 1
fi

export GOROOT=${GOROOT:-$(go env GOROOT)}

make build

"$RENODE" --disable-xwt \
  -e 'mach create "stm32l432-keyboard"' \
  -e 'machine LoadPlatformDescription @renode/stm32l432_keyboard.repl' \
  -e "sysbus LoadELF @$FIRMWARE" \
  -e 'start' \
  2>&1 | tee "$RENODE_LOG"
