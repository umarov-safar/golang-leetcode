package wordbreak

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}

	}

	return dp[n]
}
