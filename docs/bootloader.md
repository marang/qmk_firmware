# Bootloader Recovery

The `flash` make target builds the firmware with TinyGo and uploads it to the
board using the STM32 DFU protocol. If the upload fails and the board no longer
enumerates over USB, force the MCU into its system bootloader:

1. Disconnect USB.
2. Hold the `BOOT0` button or short the `BOOT0` pin to 3.3V.
3. Press and release `RESET`.
4. Release `BOOT0` and re-connect the board.
5. Run `make flash` again to upload a working image.

If the device still fails to program, erase the flash with:

```sh
 dfu-util -a 0 -e
```

Then retry `make flash`.
