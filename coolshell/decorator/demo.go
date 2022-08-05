package main_1

import "fmt"

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("decorator start")
		f(s)
		fmt.Println("decoraor end")
	}
}

func hello(s string) {
	fmt.Println(s)
}

func main() {
	decorator(hello)("cjp")
}
