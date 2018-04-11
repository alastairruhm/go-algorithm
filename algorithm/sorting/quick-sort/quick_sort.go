package quicksort

// QuickSort ...
func QuickSort(array []int) []int {
	var sorted []int
	if len(array) < 2 { // array is sorted if the length is 0 or 1,
		sorted = array
		// return array
	} else {
		pivot := array[0] // pivot value use to compare
		var less []int
		var greater []int
		for _, e := range array[1:] {
			if e <= pivot {
				less = append(less, e)
			} else {
				greater = append(greater, e)
			}
		}
		sorted = append(sorted, QuickSort(less)...)
		sorted = append(sorted, pivot)
		sorted = append(sorted, QuickSort(greater)...)
		// return sorted
	}
	return sorted
}
