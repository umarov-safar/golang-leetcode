package main

func findMiddleIndex(nums []int) int {
	s := findSum(nums)

	l := 0
	left := 0
	for l < len(nums) {
		right := s - left - nums[l]

		if right == left {
			return l
		}

		left += nums[l]
		l++
	}

	return -1
}

func findSum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}

	return s
}
