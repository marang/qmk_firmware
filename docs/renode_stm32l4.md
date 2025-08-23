# Running STM32L4 Firmware in Renode

A helper script is provided to launch [Renode](https://renode.io) with the STM32L4 platform description.

## Usage

```
util/renode_stm32l4.sh <path/to/elf>
```

The script checks that `renode` is available in your `PATH` and starts the emulator with the specified firmware image.
