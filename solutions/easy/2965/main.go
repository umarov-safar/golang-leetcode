package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	m := make(map[int]int)
	for _, raw := range grid {
		for _, v := range raw {
			m[v]++
		}
	}

	x, y := 0, 0
	for i := 1; i <= len(m)+1; i++ {
		v, ok := m[i]
		if ok && v > 1 {
			x = i
		} else if !ok {
			y = i
		}

		if x > 0 && y > 0 {
			return []int{x, y}
		}
	}

	return []int{}
}
