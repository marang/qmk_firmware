#!/bin/bash
# Debug script to test Renode step by step

set -euo pipefail

echo "=== DEBUG RENODE TEST ==="

# First, let's make sure our firmware builds
echo "Building firmware..."
make build

echo "Firmware built successfully"

# Test 1: Simple machine creation and firmware loading
echo "Test 1: Basic firmware loading"
timeout 5s renode --disable-xwt \
  -e 'mach create "test"' \
  -e 'sysbus LoadELF @build/firmware.elf' \
  -e 'quit' \
  2>&1 | head -20

echo ""
echo "Test 2: Full platform with firmware loading"
timeout 10s renode --disable-xwt \
  -e 'mach create "test"' \
  -e 'machine LoadPlatformDescription @renode/stm32l432_keyboard.repl' \
  -e "sysbus LoadELF @build/firmware.elf" \
  -e 'quit' \
  2>&1 | head -20

echo ""
echo "Test 3: Try without quit to see if it hangs"
timeout 3s renode --disable-xwt \
  -e 'mach create "test"' \
  -e 'machine LoadPlatformDescription @renode/stm32l432_keyboard.repl' \
  -e "sysbus LoadELF @build/firmware.elf" \
  -e 'cpu SingleStep' \
  -e 'cpu SingleStep' \
  2>&1 | head -10

echo "=== END DEBUG ==="