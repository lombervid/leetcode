package twosum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v: %v", test.nums, test.target)
		t.Run(testName, func(t *testing.T) {
			got := twoSum(test.nums, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
