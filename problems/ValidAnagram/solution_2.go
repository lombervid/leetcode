package validanagram

import "strings"

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	for _, char := range s {
		if index := strings.Index(t, string(char)); index > -1 {
			t = t[:index] + t[index+1:]
			continue
		}
		return false
	}

	return true
}
