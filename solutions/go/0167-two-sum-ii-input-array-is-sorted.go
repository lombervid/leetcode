package main

func twoSum(numbers []int, target int) []int {
	pairs := make(map[int]int)

	for key, value := range numbers {
		if pair, ok := pairs[target-value]; ok {
			return []int{pair, key + 1}
		}

		pairs[value] = key + 1
	}

	return nil
}
