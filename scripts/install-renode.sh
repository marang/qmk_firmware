#!/bin/bash
set -euo pipefail

RENODE_URL="https://builds.renode.io/renode-latest.linux-portable.tar.gz"
INSTALL_DIR="${RENODE_INSTALL_DIR:-$HOME/renode}"

mkdir -p "$INSTALL_DIR"
curl -L "$RENODE_URL" | tar -xz -C "$INSTALL_DIR" --strip-components=1

# Make renode available on PATH
if [ -n "${GITHUB_PATH:-}" ]; then
  echo "$INSTALL_DIR" >> "$GITHUB_PATH"
fi
export PATH="$INSTALL_DIR:$PATH"

echo "Renode installed to $INSTALL_DIR"
