package selectionsort

// SelectionSort ...
func SelectionSort(array []int) []int {
	var newArr []int
	for range array {
		smallest := findSmallest(array)
		newArr = append(newArr, array[smallest])
		array = append(array[:smallest], array[smallest+1:]...)
	}
	return newArr
}

// findSmallest finds out the index of smallest element
func findSmallest(array []int) int {
	smallest := array[0]
	index := 0
	for i := 1; i < len(array); i++ {
		if array[i] < smallest {
			smallest = array[i]
			index = i
		}
	}
	return index
}
