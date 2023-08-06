package binarysearch

import (
	"sync"
)

// reimplemented binary search to reflect the usage of concurrency
func binarySearch(data []int, value int, result *bool) {
	low := 0
	high := len(data) - 1

	for low <= high {
		median := (low + high) / 2

		if data[median] < value {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(data) || data[low] != value {
		return
	}

	*result = true
}

func ConcurrentSearch(data []int, value int) bool {
	result := false
	wg := sync.WaitGroup{}

	for i := 0; i < len(data); i += 3 {
		wg.Add(1)

		if i+3 >= len(data) {
			go func(i int) {
				binarySearch(data[i:], value, &result)
				wg.Done()
			}(i)
		} else {
			go func(i int) {
				binarySearch(data[i:i+3], value, &result)
				wg.Done()
			}(i)
		}

		if i+3 >= len(data) {
			break
		}
	}

	wg.Wait()

	return result
}
