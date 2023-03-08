package main

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		word := []rune(words[i])

		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			word[i], word[j] = word[j], word[i]
		}

		words[i] = string(word)
	}

	return strings.Join(words, " ")
}
