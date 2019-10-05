package tzpure

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func u64() uint64 {
	return rand.Uint64() & (math.MaxUint64 >> 1)
}

func TestSL2_MarshalBinary(t *testing.T) {
	a := new(SL2)
	a[0][0] = *NewGF127(u64(), u64())
	a[0][1] = *NewGF127(u64(), u64())
	a[1][0] = *NewGF127(u64(), u64())
	a[1][1] = *NewGF127(u64(), u64())

	data, err := a.MarshalBinary()
	require.NoError(t, err)

	b := new(SL2)
	err = b.UnmarshalBinary(data)
	require.NoError(t, err)

	require.Equal(t, a, b)
}
