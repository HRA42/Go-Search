package binarysearch

func Search(index []int, target int) bool {
	low := 0
	high := len(index) - 1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case index[mid] < target:
			low = mid + 1
		case index[mid] > target:
			high = mid - 1
		default:
			return true
		}
	}
	return false
}
