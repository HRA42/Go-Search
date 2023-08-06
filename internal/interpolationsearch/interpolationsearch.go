package interpolationsearch

func Search(arr []int, target int) bool {
	start := 0
	end := len(arr) - 1

	for start <= end && target >= arr[start] && target <= arr[end] {
		position := start + (((end - start) / (arr[end] - arr[start])) * (target - arr[start]))

		if arr[position] == target {
			return true
		} else if arr[position] < target {
			start = position + 1
		} else {
			end = position - 1
		}
	}
	return false
}
