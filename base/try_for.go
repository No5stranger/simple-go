package base

import "fmt"

func TryFora() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				//break
				continue
			}
			fmt.Println("hello")
		}
		fmt.Println("world")
	}
}
