# QMK Firmware (TinyGo)

This repository contains a TinyGo-based firmware for the Keychron Q1 Pro keyboard and supporting packages.

## Build

1. Install [TinyGo](https://tinygo.org/getting-started/) and ensure it can target the STM32L432 MCU.
2. Build the firmware with:
   ```sh
   make build
   ```
   The compiled ELF image will be placed in `build/firmware.elf`.

## Flash

Upload the firmware to hardware using the STM32 DFU bootloader:

```sh
make flash
```

If flashing fails, consult [docs/bootloader.md](docs/bootloader.md) for recovery steps.

## Test

Integration tests run inside [Renode](https://renode.io/). Execute the Renode harness after building:

```sh
scripts/run-renode.sh
```

Logs are captured to `renode.log` for analysis.

## Quick Start

1. Install TinyGo and Renode.
2. Build the firmware: `make build`
3. Run the Renode test setup: `scripts/run-renode.sh`
4. Flash to hardware when ready: `make flash`

For additional debugging information and platform setup, see the documentation under `docs/`.
