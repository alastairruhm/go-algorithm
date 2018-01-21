package bs

// BinarySearch ...
func BinarySearch(sortedArray []int, item int) int {
	low, high := 0, len(sortedArray)-1

	for low <= high {
		mid := (low + high) / 2
		guess := sortedArray[mid]
		if guess == item {
			return mid
		}
		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
