package minimumdistancebetweenthreeequalelementsi

/**
3741. Minimum Distance Between Three Equal Elements II
link: https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii
*/

import "math"

func minimumDistance(nums []int) int {
	m := make(map[int][]int)

	for i, num := range nums {
		m[num] = append(m[num], i)
	}

	min := math.MaxInt
	for _, indices := range m {
		if len(indices) >= 3 {
			find := findMinDistince(indices)
			if find < min {
				min = find
			}
		}
	}

	if min == math.MaxInt {
		return -1
	}

	return min
}

func findMinDistince(indices []int) int {
	m := math.MaxInt
	for i := 0; i+2 < len(indices); i++ {
		a, j, k := indices[i], indices[i+1], indices[i+2]
		r := abs(a-j) + abs(j-k) + abs(k-a)
		if r < m {
			m = r
		}
	}

	return m
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
