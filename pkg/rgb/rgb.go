package rgb

// LED represents a single RGB LED with 8-bit color channels.
type LED struct {
	R, G, B uint8
}

// Strip manages a collection of LEDs with global brightness control.
type Strip struct {
	leds       []LED
	brightness uint8
}

// New creates a Strip containing count LEDs and enables them.
func New(count int) *Strip {
	return &Strip{
		leds:       make([]LED, count),
		brightness: 255,
	}
}

// LED returns a pointer to the LED at index i or nil if out of range.
func (s *Strip) LED(i int) *LED {
	if i < 0 || i >= len(s.leds) {
		return nil
	}
	return &s.leds[i]
}

// SetBrightness adjusts the global brightness scaling applied in Frame.
func (s *Strip) SetBrightness(b uint8) {
	s.brightness = b
}

// Frame returns a copy of the LED colors scaled by the current brightness.
func (s *Strip) Frame() []LED {
	scaled := make([]LED, len(s.leds))
	for i, l := range s.leds {
		scaled[i] = LED{
			R: scale(l.R, s.brightness),
			G: scale(l.G, s.brightness),
			B: scale(l.B, s.brightness),
		}
	}
	return scaled
}

func scale(v, b uint8) uint8 {
	return uint8(uint16(v) * uint16(b) / 255)
}
