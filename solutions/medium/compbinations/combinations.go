package compbinations

// 77. Combinations
// Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

func combine(n int, k int) [][]int {
	var res [][]int
	var dfs func(i int, path []int)

	dfs = func(i int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		if i > n {
			return
		}

		path = append(path, i)
		dfs(i+1, path)
		path = path[:len(path)-1]
		dfs(i+1, path)
	}

	dfs(1, []int{})

	return res
}
