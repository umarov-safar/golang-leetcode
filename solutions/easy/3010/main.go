package main

import "sort"

func minimumCost(nums []int) int {
	sort.Ints(nums[1:])

	return nums[0] + nums[1] + nums[2]
}

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}

	return s
}
