package tzC

//#cgo CFLAGS: -march=corei7-avx
//#cgo LDFLAGS: -L$SRCDIR/tzhash/tz/core
//#include "sl2-inl.h"
import "C"
import (
	"math"
	"unsafe"

	"github.com/nspcc-dev/tzhash/gf127"
)

var minmax = [2]gf127.GF127{{0, 0}, {math.MaxUint64, math.MaxUint64}}

type digestL struct {
	s sl2_t
}

// New returns a new hash.Hash computing the Tillich-Zémor checksum.
func NewWithGoLoop() *digestL {
	d := new(digestL)
	d.Reset()
	return d
}

func (d *digestL) Sum(in []byte) []byte {
	// Make a copy of d so that caller can keep writing and summing.
	d0 := *d
	h := d0.checkSum()
	return append(in, h[:]...)
}

func (d *digestL) checkSum() [hashSize]byte {
	return decode(d.s)
}

func (d *digestL) Reset() {
	d.s = unit()
}

func (d *digestL) Write(data []byte) (n int, err error) {
	n = len(data)
	start := uintptr(unsafe.Pointer(&d.s[0]))
	addr := []*C.gf2p127_t{
		&d.s[0],
		(*C.gf2p127_t)(unsafe.Pointer(start + uintptr(16))),
		(*C.gf2p127_t)(unsafe.Pointer(start + uintptr(32))),
		(*C.gf2p127_t)(unsafe.Pointer(start + uintptr(48))),
	}
	for _, b := range data {
		C.sl2_mul_bits_right(addr[0], addr[1], addr[2], addr[3], C.uchar(b))
	}
	return
}

func (d *digestL) Size() int {
	return hashSize
}

func (d *digestL) BlockSize() int {
	return hashBlockSize
}
