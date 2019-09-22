package base

import (
	"fmt"
)

type A struct {
	m map[int]int
}

func NewA() *A {
	a := A{}
	a.m = make(map[int]int)
	return &a
}

func (a *A) d() {
	a.m[1] = 2
	a.p()
}

func (a *A) p() {
	for k, v := range a.m {
		fmt.Println(k)
		fmt.Println(v)
	}
}
