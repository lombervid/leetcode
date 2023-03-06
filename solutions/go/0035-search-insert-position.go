package main

func searchInsert(nums []int, target int) int {
	index := len(nums)
	left, right := 0, index-1

	for left <= right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return index
}
