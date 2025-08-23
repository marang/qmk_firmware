#!/bin/bash
# Launch an STM32F4 QMK .elf under QEMU with UART logging
set -e
ELF="$1"
if [ -z "$ELF" ]; then
  echo "Usage: $0 path/to/elf" >&2
  exit 1
fi
if ! command -v qemu-system-arm >/dev/null 2>&1; then
  echo "qemu-system-arm not found" >&2
  exit 1
fi
QEMU_BIN="${QEMU_BIN:-qemu-system-arm}"
QEMU_MACHINE="${QEMU_MACHINE:-netduinoplus2}"
exec "$QEMU_BIN" -M "$QEMU_MACHINE" -kernel "$ELF" -nographic -serial stdio -monitor none -d guest_errors "$@"
