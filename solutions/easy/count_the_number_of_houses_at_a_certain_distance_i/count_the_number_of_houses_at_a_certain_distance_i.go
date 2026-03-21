package main

import "sort"

// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance/

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
