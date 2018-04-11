package quicksort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuickSort(t *testing.T) {
	var cases = []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 1, 4, 3}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 5, 2, 4, 3}, []int{1, 2, 3, 4, 5}},
	}
	// Only pass t into top-level Convey calls
	Convey("Quick Sort", t, func() {
		Convey("ok", func() {
			for _, c := range cases {
				o := QuickSort(c.input)
				So(o, ShouldResemble, c.expected)
			}
		})
	})
}
