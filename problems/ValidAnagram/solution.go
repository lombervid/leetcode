package validanagram

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
