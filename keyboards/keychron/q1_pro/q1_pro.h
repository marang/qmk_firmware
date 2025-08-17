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

#include "quantum.h"
#ifdef VIA_ENABLE
#    include "via.h"
#endif

#define ___ KC_NO

#ifdef VIA_ENABLE
#    define USER_START QK_KB_0
#else
#    define USER_START SAFE_RANGE
#endif

// clang-format off
enum {
    KC_LOPTN = USER_START,
    KC_ROPTN,
    KC_LCMMD,
    KC_RCMMD,
    KC_TASK,
    KC_FILE,
    KC_SNAP,
    KC_CTANA,
    KC_SIRI,
#ifdef KC_BLUETOOTH_ENABLE
    BT_HST1,
    BT_HST2,
    BT_HST3,
    BAT_LVL,
#else
    BT_HST1 = KC_TRNS,
    BT_HST2 = KC_TRNS,
    BT_HST3 = KC_TRNS,
    BAT_LVL = KC_TRNS,
#endif
        NEW_SAFE_RANGE
};

#define LAYOUT_ansi_82( \
    K000, K001, K002, K003, K004, K005, K006, K007, K008, K009, K00A, K00B, K00C, K00D, K00F, \
    K100, K101, K102, K103, K104, K105, K106, K107, K108, K109, K10A, K10B, K10C, K10D, K10F, \
    K200, K201, K202, K203, K204, K205, K206, K207, K208, K209, K20A, K20B, K20C, K20D, K20F, \
    K300, K301, K302, K303, K304, K305, K306, K307, K308, K309, K30A, K30B, K30D, K30F, \
    K400, K402, K403, K404, K405, K406, K407, K408, K409, K40A, K40B, K40D, K40E, \
    K500, K501, K502, K506, K50A, K50B, K50C, K50D, K50E, K50F \
) { \
    { K000, K001, K002, K003, K004, K005, K006, K007, K008, K009, K00A, K00B, K00C, K00D, ___, K00F }, \
    { K100, K101, K102, K103, K104, K105, K106, K107, K108, K109, K10A, K10B, K10C, K10D, ___, K10F }, \
    { K200, K201, K202, K203, K204, K205, K206, K207, K208, K209, K20A, K20B, K20C, K20D, ___, K20F }, \
    { K300, K301, K302, K303, K304, K305, K306, K307, K308, K309, K30A, K30B, ___, K30D, ___, K30F }, \
    { K400, ___, K402, K403, K404, K405, K406, K407, K408, K409, K40A, K40B, ___, K40D, K40E, ___ }, \
    { K500, K501, K502, ___, ___, ___, K506, ___, ___, ___, K50A, K50B, K50C, K50D, K50E, K50F } \
}

#define LAYOUT LAYOUT_ansi_82
#define LAYOUT_all LAYOUT_ansi_82
