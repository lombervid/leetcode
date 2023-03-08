package main

// Another (less efficient) solution
func moveZeroes2(nums []int) {
	var i int
	var zeroes int

	for i < len(nums)-zeroes {
		if nums[i] != 0 {
			i++
			continue
		}

		zeroes++
		for j := i + 1; j < len(nums); j++ {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func moveZeroes(nums []int) {
	var j int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		j++
	}
}
