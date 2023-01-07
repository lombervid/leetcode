package topkfrequentelements

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) == k {
		return nums
	}

	counts := make(map[int]int)
	numbers := make([]int, 0)
	for _, value := range nums {
		counts[value]++
		if counts[value] == 1 {
			numbers = append(numbers, value)
		}
	}

	sort.Slice(numbers, func(i, j int) bool {
		return counts[numbers[i]] > counts[numbers[j]]
	})

	return numbers[:k]
}
