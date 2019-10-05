package tz

import (
	"errors"
	"fmt"

	"git.nspcc.ru/nspcc/tzhash/gf127"
)

type SL2 [2][2]gf127.GF127

func (c *SL2) MarshalBinary() (data []byte, err error) {
	s := c.ByteArray()
	return s[:], nil
}

func (c *SL2) UnmarshalBinary(data []byte) (err error) {
	if len(data) != 64 {
		return errors.New("data must be 64-bytes long")
	}

	if err = c[0][0].UnmarshalBinary(data[:16]); err != nil {
		return err
	}
	if err = c[0][1].UnmarshalBinary(data[16:32]); err != nil {
		return err
	}
	if err = c[1][0].UnmarshalBinary(data[32:48]); err != nil {
		return err
	}
	if err = c[1][1].UnmarshalBinary(data[48:64]); err != nil {
		return err
	}

	return nil
}

func (c *SL2) mulStrassen(a, b *SL2, x *[8]gf127.GF127) *SL2 {
	// strassen algorithm
	gf127.Add(&a[0][0], &a[1][1], &x[0])
	gf127.Add(&b[0][0], &b[1][1], &x[1])
	gf127.Mul(&x[0], &x[1], &x[0])

	gf127.Add(&a[1][0], &a[1][1], &x[1])
	gf127.Mul(&x[1], &b[0][0], &x[1])

	gf127.Add(&b[0][1], &b[1][1], &x[2])
	gf127.Mul(&x[2], &a[0][0], &x[2])

	gf127.Add(&b[1][0], &b[0][0], &x[3])
	gf127.Mul(&x[3], &a[1][1], &x[3])

	gf127.Add(&a[0][0], &a[0][1], &x[4])
	gf127.Mul(&x[4], &b[1][1], &x[4])

	gf127.Add(&a[1][0], &a[0][0], &x[5])
	gf127.Add(&b[0][0], &b[0][1], &x[6])
	gf127.Mul(&x[5], &x[6], &x[5])

	gf127.Add(&a[0][1], &a[1][1], &x[6])
	gf127.Add(&b[1][0], &b[1][1], &x[7])
	gf127.Mul(&x[6], &x[7], &x[6])

	gf127.Add(&x[2], &x[4], &c[0][1])
	gf127.Add(&x[1], &x[3], &c[1][0])

	gf127.Add(&x[4], &x[6], &x[4])
	gf127.Add(&x[0], &x[3], &c[0][0])
	gf127.Add(&c[0][0], &x[4], &c[0][0])

	gf127.Add(&x[0], &x[1], &x[0])
	gf127.Add(&x[2], &x[5], &c[1][1])
	gf127.Add(&c[1][1], &x[0], &c[1][1])

	return c
}

func (c *SL2) mul(a, b *SL2, x *[4]gf127.GF127) *SL2 {
	// naive implementation
	gf127.Mul(&a[0][0], &b[0][0], &x[0])
	gf127.Mul(&a[0][0], &b[0][1], &x[1])
	gf127.Mul(&a[1][0], &b[0][0], &x[2])
	gf127.Mul(&a[1][0], &b[0][1], &x[3])

	gf127.Mul(&a[0][1], &b[1][0], &c[0][0])
	gf127.Add(&c[0][0], &x[0], &c[0][0])
	gf127.Mul(&a[0][1], &b[1][1], &c[0][1])
	gf127.Add(&c[0][1], &x[1], &c[0][1])
	gf127.Mul(&a[1][1], &b[1][0], &c[1][0])
	gf127.Add(&c[1][0], &x[2], &c[1][0])
	gf127.Mul(&a[1][1], &b[1][1], &c[1][1])
	gf127.Add(&c[1][1], &x[3], &c[1][1])

	return c
}

func (c *SL2) MulA() *SL2 {
	var a gf127.GF127

	gf127.Mul10(&c[0][0], &a)
	gf127.Mul1(&c[0][0], &c[0][1])
	gf127.Add(&a, &c[0][1], &c[0][0])

	gf127.Mul10(&c[1][0], &a)
	gf127.Mul1(&c[1][0], &c[1][1])
	gf127.Add(&a, &c[1][1], &c[1][0])

	return c
}

func (c *SL2) MulB() *SL2 {
	var a gf127.GF127

	gf127.Mul1(&c[0][0], &a)
	gf127.Mul10(&c[0][0], &c[0][0])
	gf127.Add(&c[0][1], &c[0][0], &c[0][0])
	gf127.Add(&c[0][0], &a, &c[0][1])

	gf127.Mul1(&c[1][0], &a)
	gf127.Mul10(&c[1][0], &c[1][0])
	gf127.Add(&c[1][1], &c[1][0], &c[1][0])
	gf127.Add(&c[1][0], &a, &c[1][1])

	return c
}

func (c *SL2) Mul(a, b *SL2) *SL2 {
	d := new([4]gf127.GF127)
	return c.mul(a, b, d)
}

func (a *SL2) String() string {
	return fmt.Sprintf("%s%s%s%s", a[0][0].String(), a[0][1].String(), a[1][0].String(), a[1][1].String())
}

func (a *SL2) ByteArray() (b [hashSize]byte) {
	t := a[0][0].ByteArray()
	copy(b[:], t[:])

	t = a[0][1].ByteArray()
	copy(b[16:], t[:])

	t = a[1][0].ByteArray()
	copy(b[32:], t[:])

	t = a[1][1].ByteArray()
	copy(b[48:], t[:])

	return
}
