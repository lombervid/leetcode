package main

/* func isAnagram2(s string, t string) bool {
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
} */

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int, len(s))

	for _, char := range s {
		counts[char]++
	}

	for _, char := range t {
		if counts[char]--; counts[char] < 0 {
			return false
		}
	}

	return true
}
