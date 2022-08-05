package main

import (
	"fmt"
)

func NewInt(before int, t int, s int) int {
	a := 1 << t
	b := ^a
	c := before & b
	d := s << t
	e := before | d
	fmt.Println(a, b, c, d, e)
	if s == 0 {
		return before & ^(1 << t)
	}
	return before | s<<t
}

func main() {
	seed := [][]int{{0, 0, 0}, {0, 0, 1}, {1, 0, 0}, {1, 1, 1}, {1, 1, 0}}
	for _, s := range seed {
		n := NewInt(s[0], s[1], s[2])
		nn := n >> s[1]
		nnn := nn & 1
		fmt.Println(s, n, nn, nnn)
	}
}
