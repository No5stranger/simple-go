package simple_channel

import "fmt"

func SendReceive() {
	ch := make(chan int, 2)
	ch2 := make(chan int, 3)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("sender send %d\n", i)
			ch <- i
		}
		fmt.Printf("sender closing channel\n")
		close(ch)
	}()
	go func() {
		for j := 0; j < 5; j++ {
			ch2 <- j
		}
		close(ch2)
	}()
	select {
	case e1, ok := <-ch:
		if ok {
			fmt.Printf("receiver 1 get %d\n", e1)
		} else {
			fmt.Printf("1 no elemement received\n")
		}
	case e2 := <-ch2:
		fmt.Printf("receiver 2 get %d\n", e2)
	}
	//for {
	//elem, ok := <-ch
	//if !ok {
	//fmt.Printf("no elemement received\n")
	//break
	//}
	//fmt.Printf("receiver get %d\n", elem)
	//}
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
