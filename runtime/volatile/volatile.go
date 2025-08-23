package volatile

type Register32 struct {
	Reg uint32
}

func (r *Register32) Get() uint32               { return r.Reg }
func (r *Register32) Set(value uint32)          { r.Reg = value }
func (r *Register32) SetBits(value uint32)      { r.Reg |= value }
func (r *Register32) ClearBits(value uint32)    { r.Reg &^= value }
func (r *Register32) HasBits(value uint32) bool { return r.Reg&value != 0 }
func (r *Register32) ReplaceBits(value, mask uint32, pos uint8) {
	r.Reg = (r.Reg & ^(mask << pos)) | (value << pos)
}
