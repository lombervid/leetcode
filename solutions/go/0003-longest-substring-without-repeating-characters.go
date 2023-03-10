package main

func lengthOfLongestSubstring(s string) int {
	longest, strlen, last := 0, 0, 0
	positions := make(map[rune]int)

	for i, char := range s {
		if prev, ok := positions[char]; ok {
			if strlen > longest {
				longest = strlen
			}
			if prev > last {
				last = prev
			}
			strlen = (i - last) - 1
		}
		positions[char] = i
		strlen++
	}

	if strlen > longest {
		return strlen
	}

	return longest
}
