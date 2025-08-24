#!/usr/bin/env bash
set -euo pipefail

if ! docker info >/dev/null 2>&1; then
  echo "Docker daemon is not reachable. Please start Docker." >&2
  exit 1
fi

act -W .github/workflows/tinygo-renode.yml -j build -P ubuntu-latest=ghcr.io/catthehacker/ubuntu:full-22.04

if [[ ! -f hid.log || ! -f build/firmware.elf ]]; then
  echo "Expected artifacts hid.log and build/firmware.elf were not produced." >&2
  exit 1
fi

echo "Workflow completed and required artifacts are present."
