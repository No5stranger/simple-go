package main

import "fmt"

func makeNew() {
	s := make([]int, 3)
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("slice len:%d cap:%d", len(s), cap(s))
		fmt.Println(s)
	}
}

func useArray() {
	var l []int = []int{1, 2, 3, 4, 5, 6}
	s := l[2:4]
	fmt.Println(s)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("slice len:%d cap:%d\n", len(s), cap(s))
		fmt.Println(s)
		fmt.Println(l)
	}
}

func main() {
	fmt.Println("test new by make\n")
	makeNew()
	fmt.Println("test new by array\n")
	useArray()
}
