package tz

import (
	"git.nspcc.ru/nspcc/tzhash/gf127"
)

var one = gf127.New(1, 0)
var zero = gf127.New(0, 0)
var x = gf127.New(2, 0)
var id = &SL2{{*one, *zero}, {*zero, *one}}

// a(b) are the matrices used as images of
// bits 0(1) in Tillich-Zémor hashing
var a = &SL2{{*x, *one}, {*one, *zero}}
var b = &SL2{{*x, *x.Add(x, one)}, {*one, *one}}

var cache [256]*SL2

func init() {
	for i := 0; i < 256; i++ {
		cache[i] = HashByte(byte(i))
	}
}

func transformByte(x byte) *SL2 {
	if x == 0 {
		return a
	}
	return b
}

// HashByte calculates hash of one byte
func HashByte(x byte) (r *SL2) {
	r = new(SL2)
	*r = *transformByte(x & 0x80)
	for i := uint8(1); i < 8; i++ {
		r.Mul(r, transformByte((x & (1 << (7 - i)))))
	}
	return r
}

// Hash calculates hash of byte-slice using homomorphic
// property of Tillich-Zémor hashing
func Hash(xs []byte) (r *SL2) {
	if len(xs) == 0 {
		return nil
	}

	r = unit()
	d := new([8]gf127.GF127)
	for _, x := range xs {
		r.mulStrassen(r, cache[x], d)
	}
	return
}
