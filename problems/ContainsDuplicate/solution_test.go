package containsduplicate

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v", test.input)
		t.Run(testName, func(t *testing.T) {
			got := containsDuplicate(test.input)
			if got != test.want {
				t.Errorf("got %t want %t", got, test.want)
			}
		})
	}
}
