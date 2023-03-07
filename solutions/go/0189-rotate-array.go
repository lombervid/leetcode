package main

// Another solution using append()
func rotate2(nums []int, k int) {
	division := len(nums) - k%len(nums)
	copy(nums, append(nums[division:], nums[:division]...))
}

func rotate(nums []int, k int) {
	k %= len(nums)

	if k == 0 {
		return
	}

	count := len(nums)
	temp := make([]int, count)
	copy(temp, nums)

	for i := 0; i < count; i++ {
		j := (count - k + i) % count
		nums[i] = temp[j]
	}
}
