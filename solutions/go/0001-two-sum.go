package main

func twoSum(nums []int, target int) []int {
	rest := make(map[int]int, 0)

	for i, num := range nums {
		if j, exists := rest[target-num]; exists {
			return []int{j, i}
		}
		rest[num] = i
	}

	return nil
}
