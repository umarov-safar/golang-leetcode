package numberofislands

// Title: 200. Number of Islands
// Link: https://leetcode.com/problems/number-of-islands/description/

func numIslands(grid [][]byte) int {
	islands := 0
	rows, cols := len(grid), len(grid[0])

	visited := make([][]bool, rows)
	for i := range rows {
		visited[i] = make([]bool, cols)
	}

	var dfs func(i, j int)

	dfs = func(i, j int) {
		queue := [][2]int{{i, j}}
		visited[i][j] = true
		directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				in, jn := curr[0]+dir[0], curr[1]+dir[1]
				if in < 0 || jn < 0 || jn >= cols || in >= rows || visited[in][jn] || grid[in][jn] != '1' {
					continue
				}

				visited[in][jn] = true

				queue = append(queue, [2]int{in, jn})
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				islands++
				dfs(i, j)
			}
		}
	}

	return islands
}
