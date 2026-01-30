package main

func countPartitions(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		left := nums[i]
		right := nums[n-1] - nums[i]
		if (left-right)%2 == 0 {
			ans++
		}
	}

	return ans
}
