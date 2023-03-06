package main

func sortedSquares(nums []int) []int {
	squares := make([]int, len(nums))
	left, right := 0, len(nums)-1
	index := right

	for index >= 0 {
		lSqrt := nums[left] * nums[left]
		rSqrt := nums[right] * nums[right]

		if lSqrt > rSqrt {
			squares[index] = lSqrt
			left++
		} else {
			squares[index] = rSqrt
			right--
		}
		index--
	}

	return squares
}
