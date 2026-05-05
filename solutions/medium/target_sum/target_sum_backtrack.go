package targetsum

// Title: 494. Target Sum
// Link: https://leetcode.com/problems/target-sum/description/

func findTargetSumWays(nums []int, target int) int {
	var backtrack func(i int, s int)
	count := 0
	n := len(nums)
	backtrack = func(i int, s int) {
		if i == n {
			if s == target {
				count++
			}
			return
		}

		backtrack(i+1, s+nums[i])
		backtrack(i+1, s-nums[i])
	}

	backtrack(0, 0)
	return count
}
