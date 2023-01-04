package validpalindrome

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%q", test.s)
		t.Run(testName, func(t *testing.T) {
			got := isPalindrome(test.s)
			if got != test.want {
				t.Errorf("got %t want %t", got, test.want)
			}
		})
	}
}
