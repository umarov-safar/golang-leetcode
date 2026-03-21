package main

// https://leetcode.com/problems/ant-on-the-boundary/

func returnToBoundaryCount(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	ans := 0

	for _, v := range nums {
		if v == 0 {
			ans++
		}
	}

	return ans
}
