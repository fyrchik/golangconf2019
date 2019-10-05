package tzC

//#cgo CFLAGS: -march=corei7-avx
//#cgo LDFLAGS: -L$SRCDIR/tzhash/tz/core
//#include "sl2-inl.h"
import "C"
import (
	"encoding/hex"
	"unsafe"
)

const (
	hashSize = 64

	hashBlockSize = 128
)

type sl2_t = *[2]C.gf2p127_t

type digest struct {
	s sl2_t
}

// New returns a new hash.Hash computing the Tillich-Zémor checksum.
func New() *digest {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Sum(in []byte) []byte {
	// Make a copy of d so that caller can keep writing and summing.
	d0 := *d
	h := d0.checkSum()
	return append(in, h[:]...)
}

func (d *digest) checkSum() [hashSize]byte {
	return decode(d.s)
}

func newSL2() (m sl2_t) {
	var sl2 [2][2]C.gf2p127_t
	return (sl2_t)(unsafe.Pointer(&sl2[0][0]))
}

func unit() (m sl2_t) {
	m = newSL2()
	C.sl2_unit(m)
	return
}

func sl2_mul(a, b sl2_t) sl2_t {
	var c [2][2]C.gf2p127_t
	C.sl2_mul((sl2_t)(unsafe.Pointer(&c[0][0])), a, b)
	return sl2_t(&c[0])
}

func (d *digest) Reset() {
	d.s = unit()
}

func decode(p sl2_t) (h [hashSize]byte) {
	// TODO find a better way to decode
	var shex [128]byte
	C.sl2_hex((*C.char)(unsafe.Pointer(&shex[0])), p)
	hex.Decode(h[:], shex[:])
	return
}

func (d *digest) Write(data []byte) (n int, err error) {
	n = len(data)
	if n > 0 {
		C.sl2_mul_buf_right(d.s, (*C.uchar)(&data[0]), C.size_t(n))
	}
	return
}

func (d *digest) Size() int {
	return hashSize
}

func (d *digest) BlockSize() int {
	return hashBlockSize
}

// Sum returnz Tillich-Zémor checksum of data
func Sum(data []byte) [hashSize]byte {
	p := unit()
	buf := (*C.uchar)(&data[0])
	l := C.size_t(len(data))
	if len(data) > 0 {
		C.sl2_mul_buf_right(p, buf, l)
	}
	return decode(p)
}
