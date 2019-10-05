package tz2

import (
	"hash"
	"math"

	"github.com/nspcc-dev/tzhash/gf127"
)

type digest2 struct {
	x [2]gf127.GF127x2
}

const (
	hashSize      = 64
	hashBlockSize = 128
)

var _ hash.Hash = (*digest2)(nil)
var minmax = [2]gf127.GF127{{0, 0}, {math.MaxUint64, math.MaxUint64}}

// New returns a new hash.Hash computing the Tillich-Zémor checksum.
func New() *digest2 {
	d := new(digest2)
	d.Reset()
	return d
}

func (d *digest2) Write(data []byte) (n int, err error) {
	//dd := new(digest2)
	// dd.Reset()
	// mulByteRightx2(&dd.x[0], &dd.x[1], data[0])
	// fmt.Println(dd.Sum(nil))
	//
	// dd.Reset()
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>7)&1])
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>6)&1])
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>5)&1])
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>4)&1])
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>3)&1])
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>2)&1])
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>1)&1])
	// mulBitRightx2(&dd.x[0], &dd.x[1], &minmax[(data[0]>>0)&1])
	// fmt.Println(dd.Sum(nil))

	n = len(data)
	for _, b := range data {
		mulByteRightx2(&d.x[0], &d.x[1], b)
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>7)&1])
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>6)&1])
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>5)&1])
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>4)&1])
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>3)&1])
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>2)&1])
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>1)&1])
		// mulBitRightx2(&d.x[0], &d.x[1], &minmax[(b>>0)&1])
	}
	return
}

func (d *digest2) Sum(in []byte) []byte {
	// Make a copy of d so that caller can keep writing and summing.
	d0 := *d
	h := d0.checkSum()
	return append(in, h[:]...)
}
func (d *digest2) Reset() {
	d.x[0] = gf127.GF127x2{1, 0, 0, 0}
	d.x[1] = gf127.GF127x2{0, 0, 1, 0}
}
func (d *digest2) Size() int      { return hashSize }
func (d *digest2) BlockSize() int { return hashBlockSize }
func (d *digest2) checkSum() (b [hashSize]byte) {
	// Matrix is stored transposed,
	// but we need to use order consistent with digest.
	h := d.x[0].ByteArray()
	copy(b[:], h[:16])
	copy(b[32:], h[16:])

	h = d.x[1].ByteArray()
	copy(b[16:], h[:16])
	copy(b[48:], h[16:])
	return
}

//go:noescape
func mulBitRightx2(c00c10 *gf127.GF127x2, c01c11 *gf127.GF127x2, e *gf127.GF127)

//go:noescape
func mulByteRightx2(c00c10 *gf127.GF127x2, c01c11 *gf127.GF127x2, b byte)
