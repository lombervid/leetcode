package groupsanagrams

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

func groupAnagrams2(strs []string) [][]string {
	groups := make([][]string, 0)

	for len(strs) > 0 {
		var current string
		current, strs = strs[len(strs)-1], strs[:len(strs)-1]
		anagrams := []string{current}

		for i := len(strs) - 1; i >= 0; i-- {
			if !isAnagram(current, strs[i]) {
				continue
			}

			anagrams = append(anagrams, strs[i])
			strs = append(strs[:i], strs[i+1:]...)
		}

		groups = append(groups, anagrams)
	}

	return groups
}
