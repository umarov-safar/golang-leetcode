package main

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	judges := make(map[int]int)
	trusts := make(map[int]int)

	for _, raw := range trust {
		judges[raw[1]]++
		trusts[raw[0]]++
	}

	for judge, count := range judges {
		if count == n-1 && trusts[judge] == 0 {
			return judge
		}
	}

	return -1
}
