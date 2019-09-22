package async

import (
	"fmt"
	"sync"
)

func WaitGroupFor() {
	var loop [][]int = [][]int{[]int{1, 10}, []int{11, 20}}
	wg := sync.WaitGroup{}
	wg.Add(3)
	for _, l := range loop {
		go func(ll []int) {
			for i := ll[0]; i < ll[1]; i++ {
				fmt.Println(i)
			}
			wg.Done()
		}(l)
	}
	wg.Wait()
	fmt.Println("done")
}
