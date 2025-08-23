#!/usr/bin/env bash
set -e

if ! command -v renode >/dev/null 2>&1; then
  echo "Error: renode is not installed or not in PATH" >&2
  exit 1
fi

if [ $# -ne 1 ]; then
  echo "Usage: $0 <path/to/elf>" >&2
  exit 1
fi

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
RENODE_RESC="$SCRIPT_DIR/renode/keychron_q1_pro.resc"

renode "$RENODE_RESC" "$1"
