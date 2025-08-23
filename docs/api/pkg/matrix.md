package matrix // import "github.com/qmk/qmk_firmware/pkg/matrix"

type Matrix struct{ ... }
    func New(data, clock, latch Pin, rows []Pin, cols int) *Matrix
type Pin interface{ ... }
