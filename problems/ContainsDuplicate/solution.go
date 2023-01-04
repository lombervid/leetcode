package containsduplicate

func containsDuplicate(nums []int) bool {
	duplicates := make(map[int]int)

	for _, item := range nums {
		if _, ok := duplicates[item]; ok {
			return true
		}

		duplicates[item] = 1
	}

	return false
}
