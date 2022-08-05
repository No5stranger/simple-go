package main

import (
	"fmt"
	"time"
)

func do() {
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now().Unix())
}

func main() {
	tic := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tic.C:
			do()
		}
	}
}
