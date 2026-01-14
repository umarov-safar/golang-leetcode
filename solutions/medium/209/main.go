package main

func minSubArrayLen(target int, nums []int) int {
	count := len(nums) + 1

	l, total := 0, 0

	for r := range nums {
		total += nums[r]

		for total >= target {
			count = min(r-l+1, count)
			total -= nums[l]
			l++
		}
	}

	if count > len(nums) {
		return 0
	}

	return count
}
