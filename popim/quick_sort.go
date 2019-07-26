package main

import "fmt"

func QuickSort(l []int) {
	PartitionSort(l, 0, len(l)-1)
}

func PartitionSort(l []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(l, start, end)
	partition(l, start, i-1)
	partition(l, i+1, end)
}

func partition(l []int, start, end int) int {
	var tmp int = l[end]
	var i int = start
	for j := start; j < end; j++ {
		if l[j] < tmp {
			if i != j {
				l[j], l[i] = l[i], l[j]
			}
			i++
		}
	}
	l[i], l[end] = l[end], l[i]
	return i
}

func main() {
	var l []int = []int{1, 5, 6, 3, 9, 4}
	QuickSort(l)
	fmt.Println(l)
}
