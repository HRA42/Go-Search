package sortarray

import (
	"sync"
)

const threshold = 2048

func ConcurrentSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	if len(arr) <= threshold {
		return sequentialSort(arr)
	}

	mid := len(arr) / 2

	var left, right []int
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		left = Sort(arr[:mid])
	}()

	go func() {
		defer wg.Done()
		right = Sort(arr[mid:])
	}()

	wg.Wait()

	return merge(left, right)
}

func sequentialSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := sequentialSort(arr[:mid])
	right := sequentialSort(arr[mid:])

	return merge(left, right)
}
