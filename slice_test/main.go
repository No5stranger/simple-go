package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 2, 3, 4}
	b := make([]int, 0, 2<<len(a))
	fmt.Println(2 << 50)
	fmt.Println(50 << 2)
	fmt.Println(len(b))
}
