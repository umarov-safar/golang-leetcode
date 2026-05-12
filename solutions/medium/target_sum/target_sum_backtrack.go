package targetsum

// Title: 494. Target Sum
// Link: https://leetcode.com/problems/target-sum/description/

func findTargetSumWays(nums []int, target int) int {
	var backtrack func(i int, s int) int
	n := len(nums)
	backtrack = func(i int, s int) int {
		if i == n {
			if s == target {
				return 1
			}
			return 0
		}

		return backtrack(i+1, s+nums[i]) + backtrack(i+1, s-nums[i])
	}

	
	return backtrack(0, 0)
}
