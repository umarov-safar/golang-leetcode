package numberofislands

// Title: 200. Number of Islands
// Link: https://leetcode.com/problems/number-of-islands/description/

func numIslands(grid [][]byte) int {
	visited := make(map[string]bool)
	islands := 0
	rows, cols := len(grid), len(grid[0])

	var dfs func(i, j int)

	dfs = func(i, j int) {
		queue := [][2]int{{i, j}}
		visited[key(i, j)] = true
		directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				in, jn := curr[0]+dir[0], curr[1]+dir[1]
				_, ok := visited[key(in, jn)]

				if ok || in < 0 || jn < 0 || jn >= cols || in >= rows || grid[in][jn] != '1' {
					continue
				}

				visited[key(in, jn)] = true

				queue = append(queue, [2]int{in, jn})
			}
		}
	}

	for i := range grid {
		for j := 0; j < len(grid[0]); j++ {
			_, ok := visited[key(i, j)]
			if grid[i][j] == '1' && !ok {
				islands++
				dfs(i, j)
			}
		}
	}

	return islands
}

func key(i, j int) string {
	return string(i) + "," + string(j)
}
