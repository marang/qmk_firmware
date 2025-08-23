# Simple N x M GPIO matrix model for Renode
"""Keyboard matrix peripheral for Renode.

Expose an N\u00d7M matrix to a Renode machine and provide monitor commands
``press <row> <col>`` and ``release <row> <col>`` to simulate key activity.
Columns are driven by a shift register controlled via ``ds``/``shcp``/``stcp``
lines, while rows are exposed as GPIO outputs to the SoC.
"""

from Renode import *
from Renode.Peripherals import IManagedGPIOReceiver, IProvidesGPIO
from Renode.Utilities import Log
from Renode.Peripherals.GPIOPort import GPIO


class Matrix(IManagedGPIOReceiver, IProvidesGPIO):
    """Simple keyboard matrix model with shift-register driven columns."""

    def __init__(self, machine, rows=1, cols=1):
        self.machine = machine
        self.rows = rows
        self.cols = cols
        self.keys = [[False for _ in range(cols)] for _ in range(rows)]
        # row output lines
        self.row_lines = [GPIO(f"row{i}") for i in range(rows)]

        # shift register control lines provided by the SoC
        self.ds = GPIO("ds")
        self.shcp = GPIO("shcp")
        self.stcp = GPIO("stcp")

        self._ds_state = False
        self._shift = 0

    def _update_rows(self):
        # interpret shift register as column select, active low
        for r in range(self.rows):
            level = True
            for c in range(self.cols):
                active = not ((self._shift >> c) & 1)
                if active and self.keys[r][c]:
                    level = False
                    break
            self.row_lines[r].Set(level)

    def OnGPIO(self, gpio, value):
        if gpio == self.ds:
            self._ds_state = value
        elif gpio == self.shcp and value:
            # rising edge shifts in DS value (LSB first)
            self._shift = ((self._shift << 1) | (1 if self._ds_state else 0)) & ((1 << self.cols) - 1)
        elif gpio == self.stcp and value:
            # latch triggers evaluation of row outputs
            self._update_rows()

    def Reset(self):
        self._shift = 0
        self._ds_state = False
        for r in range(self.rows):
            for c in range(self.cols):
                self.keys[r][c] = False
            self.row_lines[r].Set(True)
        self._update_rows()

    # exported commands
    @export
    def press(self, row: int, col: int):
        if 0 <= row < self.rows and 0 <= col < self.cols:
            self.keys[row][col] = True
            self._update_rows()

    @export
    def release(self, row: int, col: int):
        if 0 <= row < self.rows and 0 <= col < self.cols:
            self.keys[row][col] = False
            self._update_rows()
