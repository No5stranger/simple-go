package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")
	format := "2006-01-02 15:04:05"
	fmt.Println(time.Now().Format(format))
	rand.Seed(5)
	n := rand.Intn(20)
	fmt.Println(n)
	fmt.Println((time.Duration(n) * time.Second).Seconds())
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(time.Now().Format(format))
}
