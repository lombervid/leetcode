package main

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)

	return strings.Join(str, "")
}

func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string, 0)
	for _, word := range strs {
		key := sortString(word)
		anagramGroups[key] = append(anagramGroups[key], word)
	}

	groups := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		groups = append(groups, group)
	}

	return groups
}
