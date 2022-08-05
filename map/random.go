package main

import (
	"fmt"
)

var am map[int]string = map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
	5: "e",
	6: "f",
	7: "g",
}

func main() {
	var b []string
	for _, v := range am {
		b = append(b, v)
	}
	fmt.Printf("randon?, %s", b[:3])
}
