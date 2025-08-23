# TinyGo Setup

This guide describes how to install TinyGo and configure the Keychron Q1 Pro board.

## Installation

1. Download the TinyGo package for your platform from the [TinyGo releases](https://github.com/tinygo-org/tinygo/releases) page.
2. Install the package. For example on Debian/Ubuntu:

   ```sh
   wget https://github.com/tinygo-org/tinygo/releases/download/v0.39.0/tinygo_0.39.0_amd64.deb
   sudo dpkg -i tinygo_0.39.0_amd64.deb
   ```

3. Verify the installation and that TinyGo can target Cortex-M4 MCUs:

   ```sh
   tinygo info stm32f4disco
   ```

   The output should report an LLVM triple of `thumbv7em-unknown-unknown-eabi`.

## Building for Keychron Q1 Pro

The Keychron Q1 Pro uses an STM32L432 MCU. After adding the board definition
and machine support found in this repository, you can build firmware with:

```sh
tinygo build -target board/keychron_q1_pro ./cmd/firmware
```

