package targetsum

func findTargetSumWaysKnapsack(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if (sum+target)%2 != 0 || abs(target) > sum {
		return 0
	}

	P := (sum + target) / 2

	dp := make([]int, P+1)
	dp[0] = 1

	for _, num := range nums {
		for j := P; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[P]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
