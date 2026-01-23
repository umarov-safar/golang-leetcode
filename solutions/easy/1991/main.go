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

func findMiddleIndexPrefixSumClassic(nums []int) int {
	n := len(nums)
	ps := make([]int, n)

	i := 1
	ps[0] = nums[0]

	for i < n {
		ps[i] = nums[i] + ps[i-1]
		i++
	}

	right := ps[n-1]
	for i := 0; i < n; i++ {
		var left int
		if i == 0 {
			left = 0
		} else {
			left = ps[i-1]
		}
		if right-left-nums[i] == left {
			return i
		}
	}

	return -1
}
