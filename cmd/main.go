package main

import (
	"fmt"

	"git.nspcc.ru/nspcc/tzhash/gf127"
)

func main() {
	var a, b *gf127.GF127
	m := new(gf127.GF127)
	s := new(gf127.GF127)

	a = gf127.New(3, 0)
	b = gf127.New(0, 0xf)
	s.Add(a, b)
	m.Mul(a, b)
	fmt.Printf("%s\n%s\n%s\n%s\n\n", a.String(), b.String(), s.String(), m.String())

	// (x+1)*(x^63+x^62+...+1) = x^64-1
	a = gf127.New(3, 0)
	b = gf127.New(0xFFFFFFFFFFFFFFFF, 0)
	s.Add(a, b)
	m.Mul(a, b)
	fmt.Printf("%s\n%s\n%s\n%s\n\n", a.String(), b.String(), s.String(), m.String())

	// x^126 * x^2 == x^128 == x^64 + x
	a = gf127.New(0, 1<<62)
	b = gf127.New(4, 0)
	s.Add(a, b)
	m.Mul(a, b)
	fmt.Printf("%s\n%s\n%s\n%s\n\n", a.String(), b.String(), s.String(), m.String())

	// (x^64+x^63+1) * (x^64+x) ==
	// x^128+x^65+x^127+x^64+x^64+x ==
	// x^65+x^64+x^63+1
	a = gf127.New(1+1<<63, 1)
	b = gf127.New(2, 1)
	s.Add(a, b)
	m.Mul(a, b)
	fmt.Printf("%s\n%s\n%s\n%s\n\n", a.String(), b.String(), s.String(), m.String())
}
