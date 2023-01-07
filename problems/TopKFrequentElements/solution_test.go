package topkfrequentelements

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{4, 1, -1, 2, -1, 2, 3}, 2, []int{-1, 2}},
		{[]int{6, 0, 1, 4, 9, 7, -3, 1, -4, -8, 4, -7, -3, 3, 2, -3, 9, 5, -4, 0}, 6, []int{-3, -4, 0, 1, 4, 9}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v: %v", test.nums, test.k)
		t.Run(testName, func(t *testing.T) {
			got := topKFrequent(test.nums, test.k)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
