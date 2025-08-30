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

echo "Building firmware..."
make build

echo "Starting Renode test with minimal platform..."
echo "Firmware location: $FIRMWARE"
echo "Platform file: renode/stm32l432_keyboard.repl"

# Run Renode with verbose output for debugging
"$RENODE" --disable-xwt \
  -e 'mach create "stm32l432-keyboard"' \
  -e 'machine LoadPlatformDescription @renode/stm32l432_keyboard.repl' \
  -e "sysbus LoadELF @$FIRMWARE" \
  -e 'start' \
  2>&1 | tee "$RENODE_LOG"

echo "Renode test completed"
echo "Log file: $RENODE_LOG"
