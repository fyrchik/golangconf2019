package tzbitspure

import (
	"math"
)

type digest struct {
	x [4]GF127
}

const (
	hashSize      = 64
	hashBlockSize = 128
)

// x127x64 represents x^127 + x^63. Used in assembly file.
var x127x63 = GF127{1 << 63, 1 << 63}
var minmax = [2]GF127{{0, 0}, {math.MaxUint64, math.MaxUint64}}

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
	return d.byteArray()
}

func (d *digest) byteArray() (b [hashSize]byte) {
	for i := 0; i < 4; i++ {
		t := d.x[i].ByteArray()
		copy(b[i*16:], t[:])
	}
	return
}

func (d *digest) Reset() {
	d.x[0] = GF127{1, 0}
	d.x[1] = GF127{0, 0}
	d.x[2] = GF127{0, 0}
	d.x[3] = GF127{1, 0}
}

func (d *digest) Write(data []byte) (n int, err error) {
	n = len(data)
	for _, b := range data {
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>7)&1 != 0)
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>6)&1 != 0)
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>5)&1 != 0)
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>4)&1 != 0)
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>3)&1 != 0)
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>2)&1 != 0)
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>1)&1 != 0)
		mulBitRight(&d.x[0], &d.x[1], &d.x[2], &d.x[3], (b>>0)&1 != 0)
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
	d := New()
	d.Write(data)
	return d.checkSum()
}


func mulBitRight(c00, c01, c10, c11 *GF127, bit bool) {
	var tmp GF127
	if bit {
		tmp = *c00
		Mul10(c00, c00)
		Add(c00, c01, c00)
		Mul11(&tmp, &tmp)
		Add(c01, &tmp, c01)

		tmp = *c10
		Mul10(c10, c10)
		Add(c10, c11, c10)
		Mul11(&tmp, &tmp)
		Add(c11, &tmp, c11)
	} else {
		tmp = *c00
		Mul10(c00, c00)
		Add(c00, c01, c00)
		*c01 = tmp

		tmp = *c10
		Mul10(c10, c10)
		Add(c10, c11, c10)
		*c11 = tmp
	}
}
