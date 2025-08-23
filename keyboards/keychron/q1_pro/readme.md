# Keychron Q1 Pro

![Keychron Q1 Pro](https://cdn.shopify.com/s/files/1/0059/0630/1017/t/5/assets/keychronq1proqmkviacustommechanicalkeyboard--edited-1669962623486.jpg)

A customizable 81 keys TKL keyboard.

* Keyboard Maintainer: [Keychron](https://github.com/keychron)
* Hardware Supported: Keychron Q1 Pro
* Hardware Availability: [Keychron Q1 Pro Wireless, Customizeable, QMK/VIA Mechanical Keyboard](https://www.keychron.com/products/keychron-k8-pro-qmk-via-wireless-mechanical-keyboard)

Make example for this keyboard (after setting up your build environment):

    make keychron/q1_pro/ansi_knob:default

Flashing example for this keyboard:

    make keychron/q1_pro/ansi_knob:default:flash

**Reset Key**: Connect the USB cable, toggle mode switch to "Off", hold down the *Esc* key or reset button underneath space bar, then toggle then switch to "Cable".

See the [build environment setup](https://docs.qmk.fm/#/getting_started_build_tools) and the [make instructions](https://docs.qmk.fm/#/getting_started_make_guide) for more information. Brand new to QMK? Start with our [Complete Newbs Guide](https://docs.qmk.fm/#/newbs).

## Renode Virtual Testing

Renode 1.15 or newer can emulate the Q1 Pro so you can test keymaps without hardware.

1. Install Renode ≥1.15 (see [Renode Setup](../../../docs/renode_setup.md) for OS-specific installation steps).
2. From the repository root, build the firmware and launch Renode:

    ```sh
    make keychron/q1_pro:renode
    ```
3. At the Renode monitor, use the `matrix` peripheral to simulate key presses. For example, to press and release the *Esc* key:

    ```text
    matrix press 0 0
    matrix release 0 0
    ```

