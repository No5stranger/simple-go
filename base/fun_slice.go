package base

import (
	"fmt"
)

func FunSlice() {
	// []
	a := []int{}
	// [0]
	b := make([]int, 1)
	// [0]
	c := make([]int, 1, 3)
	printSlice(a)
	printSlice(b)
	printSlice(c)
	for i := 1; i < 5; i++ {
		a = append(a, i)
		b = append(b, i)
		c = append(c, i)
		fmt.Println("loop i:", i)
		printSlice(a)
		printSlice(b)
		printSlice(c)
	}
}

func printSlice(s []int) {
	fmt.Println(s)
}
