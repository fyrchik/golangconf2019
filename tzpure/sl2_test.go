package tzpure

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"git.nspcc.ru/nspcc/tzhash/gf127"
	. "github.com/onsi/gomega"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func u64() uint64 {
	return rand.Uint64() & (math.MaxUint64 >> 1)
}

func TestSL2_MarshalBinary(t *testing.T) {
	g := NewGomegaWithT(t)
	a := new(SL2)
	a[0][0] = *gf127.New(u64(), u64())
	a[0][1] = *gf127.New(u64(), u64())
	a[1][0] = *gf127.New(u64(), u64())
	a[1][1] = *gf127.New(u64(), u64())

	data, err := a.MarshalBinary()
	g.Expect(err).NotTo(HaveOccurred())

	b := new(SL2)
	err = b.UnmarshalBinary(data)
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(a).To(Equal(b))
}
