package simple_channel

import "fmt"

func RandomGo() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
