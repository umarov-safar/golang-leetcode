package main

// https://leetcode.com/problems/flip-square-submatrix-vertically/

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	endX := x + k - 1

	for endX > x {
		startY := y
		for startY < y+k {
			grid[endX][startY], grid[x][startY] = grid[x][startY], grid[endX][startY]
			startY++
		}
		x++
		endX--
	}

	return grid
}
