package validanagram

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%q,%q", test.s, test.t)
		t.Run(testName, func(t *testing.T) {
			got := isAnagram(test.s, test.t)
			if got != test.want {
				t.Errorf("got %t want %t", got, test.want)
			}
		})
	}
}
