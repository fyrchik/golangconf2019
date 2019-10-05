package tzpure

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func u64() uint64 {
	return rand.Uint64() & (math.MaxUint64 >> 1)
}

func TestSL2_MarshalBinary(t *testing.T) {
	// a := new(SL2)
	// a[0][0] = *NewGF127(u64(), u64())
	// a[0][1] = *NewGF127(u64(), u64())
	// a[1][0] = *NewGF127(u64(), u64())
	// a[1][1] = *NewGF127(u64(), u64())
	//
	// data, err := a.MarshalBinary()
	// g.Expect(err).NotTo(HaveOccurred())
	//
	// b := new(SL2)
	// err = b.UnmarshalBinary(data)
	// g.Expect(err).NotTo(HaveOccurred())
	//
	// g.Expect(a).To(Equal(b))
}
