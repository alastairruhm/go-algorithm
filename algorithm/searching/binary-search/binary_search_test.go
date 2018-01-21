package bs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearch(t *testing.T) {
	var cases = []struct {
		sortedArray []int
		item        int
		expected    int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
	}
	// Only pass t into top-level Convey calls
	Convey("binary search", t, func() {
		Convey("ok", func() {
			for _, c := range cases {
				o := BinarySearch(c.sortedArray, c.item)
				So(o, ShouldEqual, c.expected)
			}
		})
	})
}
