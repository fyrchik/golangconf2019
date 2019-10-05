package tzpure

import (
	"fmt"
)

const (
	hashSize = 64

	// TODO implement caching in Write so that hashBlockSize will be used
	hashBlockSize = 128
)

type digest struct {
	s SL2
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
	return d.s.ByteArray()
}

func unit() (m *SL2) {
	m = new(SL2)
	*m = *id
	return
}

func (d *digest) Dump() {
	fmt.Println(d.s)
}

func (d *digest) Reset() {
	d.s = *unit()
}

func (d *digest) Write(data []byte) (n int, err error) {
	n = len(data)
	if n > 0 {
		d.s.Mul(&d.s, Hash(data))
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
	return Hash(data).ByteArray()
}
