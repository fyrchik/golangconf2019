package tzC

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

func TestHash(t *testing.T) {
	b := make([]byte, 64)
	n, err := rand.Read(b)
	if n != 64 || err != nil {
		t.Fatal("cannot init random byte buffer")
	}

	// Test if our hashing is really homomorphic
	h := New()
	h.Write(b)
	h1 := New()
	h1.Write(b[:32])
	h2 := New()
	h2.Write(b[32:])
	h3 := decode(sl2_mul(h1.s, h2.s))

	s := hex.EncodeToString(h.Sum(nil))
	s3 := hex.EncodeToString(h3[:])
	if s != s3 {
		t.Errorf("expected (%s), got (%s)", s, s3)
	}
}
