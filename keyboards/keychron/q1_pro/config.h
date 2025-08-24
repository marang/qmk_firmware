/* Copyright 2023 @ Keychron (https://www.keychron.com)
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

#pragma once

/* Key matrix size */
#define MATRIX_ROWS 6
#define MATRIX_COLS 16

/* Turn off effects when suspended */
#define RGB_MATRIX_SLEEP

/* DIP switch for Mac/win OS switch */
#define DIP_SWITCH_PINS { A8 }

#ifdef RGB_MATRIX_ENABLE
/* RGB Matrix driver configuration */
#    define DRIVER_COUNT 2

#    define SNLED27351_I2C_ADDRESS_1 SNLED27351_I2C_ADDRESS_VDDIO
#    define SNLED27351_I2C_ADDRESS_2 SNLED27351_I2C_ADDRESS_GND
#    define SNLED27351_PHASE_CHANNEL SNLED27351_SCAN_PHASE_9_CHANNEL
#    define DRIVER_1_LED_COUNT 45
#    define DRIVER_2_LED_COUNT 37
#    define RGB_MATRIX_LED_COUNT (DRIVER_1_LED_COUNT + DRIVER_2_LED_COUNT)

/* Set to infinite, which is used in USB mode by default */
#    define RGB_MATRIX_TIMEOUT RGB_MATRIX_TIMEOUT_INFINITE

/* Allow shutdown of led driver to save power */
#    define RGB_MATRIX_DRIVER_SHUTDOWN_ENABLE
/* Turn off backlight on low brightness to save power */
#    define RGB_MATRIX_BRIGHTNESS_TURN_OFF_VAL 32

#    define CAPS_LOCK_LED_INDEX 45
#    define CAPS_LOCK_INDEX CAPS_LOCK_LED_INDEX
#    define LOW_BAT_IND_INDEX { 75 }

/* RGB Matrix Animation modes. Explicitly enabled
 * For full list of effects, see:
 * https://docs.qmk.fm/#/feature_rgb_matrix?id=rgb-matrix-effects
 */
#    define RGB_MATRIX_KEYPRESSES
#    define RGB_MATRIX_FRAMEBUFFER_EFFECTS

/* Set LED driver current */
#    define SNLED27351_CURRENT_TUNE \
        { 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38 }

#endif

#ifdef KC_BLUETOOTH_ENABLE
/* Hardware configuration */
#    define USB_BT_MODE_SELECT_PIN C15
#    define BT_MODE_SELECT_PIN USB_BT_MODE_SELECT_PIN
#    define P2P4_MODE_SELECT_PIN USB_BT_MODE_SELECT_PIN

#    define LKBT51_RESET_PIN A9
#    define BLUETOOTH_INT_OUTPUT_PIN A5
#    define LKBT51_INT_INPUT_PIN A6
#    define BLUETOOTH_INT_INPUT_PIN LKBT51_INT_INPUT_PIN
#    define SPI_SCK_PIN A5
#    define SPI_MISO_PIN A6
#    define SPI_MOSI_PIN A7

#    define USB_POWER_SENSE_PIN B1
#    define USB_POWER_CONNECTED_LEVEL 0

#    define HOST_DEVICES_COUNT 3

#    if defined(RGB_MATRIX_ENABLE)

#        define LED_DRIVER_SHUTDOWN_PIN C14

#        define HOST_LED_MATRIX_LIST \
            { 16, 17, 18 }
#        define BT_HOST_LED_MATRIX_LIST HOST_LED_MATRIX_LIST
#        define P2P4G_HOST_LED_MATRIX_LIST { 16 }

#        define BAT_LEVEL_LED_LIST \
            { 16, 17, 18, 19, 20, 21, 22, 23, 24, 25 }

/* Backlit disable timeout when keyboard is disconnected(unit: second) */
#        define DISCONNECTED_BACKLIGHT_DISABLE_TIMEOUT 40

/* Backlit disable timeout when keyboard is connected(unit: second) */
#        define CONNECTED_BACKLIGHT_DISABLE_TIMEOUT 600
#    endif

/* Keep USB connection in blueooth mode */
#    define KEEP_USB_CONNECTION_IN_BLUETOOTH_MODE

/* Enable bluetooth NKRO */
#    define BLUETOOTH_NKRO_ENABLE

/* Raw hid command for factory test and bluetooth DFU */
#    define RAW_HID_CMD 0xAA ... 0xAB
#else
/* Raw hid command for factory test */
#    define RAW_HID_CMD 0xAB
#endif

/* Emulated EEPROM configuration */
#define FEE_DENSITY_BYTES FEE_PAGE_SIZE
#define WEAR_LEVELING_BACKING_SIZE 4096
#define DYNAMIC_KEYMAP_EEPROM_MAX_ADDR 2047

/* Encoder Configuration */
#ifdef ENCODER_ENABLE
#    define ENCODER_DEFAULT_POS 0x3
#endif

/* Factory test keys */
#define FN_KEY1 MO(1)
#define FN_KEY2 MO(3)
#define BL_TEST_KEY2 KC_END

#define INVERT_OS_SWITCH_STATE

