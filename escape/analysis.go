package main

import (
	"fmt"
	"strconv"
)

func dummy(b int) int {
	var c int
	c = b
	return c
}

func incr2(u int) {
	ur := strconv.Itoa(u)
	_ = ur
}

func main() {
	var a int
	//fmt.Println(a, dummy(0))
	//incr(a)
	incr2(a)
}
