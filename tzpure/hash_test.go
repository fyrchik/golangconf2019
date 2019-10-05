package tzpure

import (
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
	h := Hash(b)
	h1 := Hash(b[:32])
	h2 := Hash(b[32:])
	if c := h1.Mul(h1, h2).String(); h.String() != c {
		t.Errorf("expected (%s), got (%s)", h.String(), c)
	}
}
