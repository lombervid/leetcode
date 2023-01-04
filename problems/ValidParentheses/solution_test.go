package validparentheses

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"]", false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%q", test.s)
		t.Run(testName, func(t *testing.T) {
			got := isValid(test.s)
			if got != test.want {
				t.Errorf("got %t want %t", got, test.want)
			}
		})
	}
}
