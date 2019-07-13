package simple_channel

import "fmt"

func SendReceive() {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("sender send %d\n", i)
		}
		fmt.Printf("sender closing channel")
		close(ch)
	}()
	for {
		elem, ok := <-ch
		if !ok {
			fmt.Printf("no elemement received\n")
			break
		}
		fmt.Printf("receiver get %d", elem)
	}
}

func RangeChannel() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	fmt.Printf("try to range channel\n")
	for eleme := range ch {
		fmt.Printf("range channel get %d\n", eleme)
	}
}
