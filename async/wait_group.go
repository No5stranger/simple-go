package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroupFor() {
	var loop [][]int = [][]int{[]int{1, 10}, []int{11, 50}}
	wg := sync.WaitGroup{}
	ch := make(chan int, 3)
	go func() {
		for {
			select {
			case c, ok := <-ch:
				if ok {
					fmt.Println(c)
					time.Sleep(500 * time.Millisecond)
					wg.Done()
				}
			}
		}
	}()
	for _, l := range loop {
		wg.Add(l[1] - l[0])
		go func(ll []int, wg *sync.WaitGroup) {
			for i := ll[0]; i < ll[1]; i++ {
				ch <- i
			}
		}(l, &wg)
	}
	wg.Wait()
	fmt.Println("done")
}

func AnotherWaitGroup() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 3)
	go func() {
		for {
			select {
			case c, ok := <-ch:
				if ok {
					fmt.Println(c)
					wg.Done()
				}
			}
		}
	}()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(m int, wg *sync.WaitGroup) {
			ch <- m
		}(i, &wg)
	}
	wg.Wait()
	fmt.Println("another done")
}

func DoTwoWaitGroup() {
	WaitGroupFor()
	AnotherWaitGroup()
}

func main() {
	DoTwoWaitGroup()
}
