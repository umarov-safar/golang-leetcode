package easy

func CheckValid(matrix [][]int) bool {
	for i, row := range matrix {
		r, c := make(map[int]bool), make(map[int]bool)

		for j, v := range row {
			if _, ok := r[v]; ok {
				return false
			} else {
				r[v] = true
			}

			if _, ok := c[matrix[j][i]]; ok {
				return false
			} else {
				c[matrix[j][i]] = true
			}
		}
	}

	return true
}

func CheckValidV2(matrix [][]int) bool {
	for i, row := range matrix {
		rowStore, colStore := [101]int{}, [101]int{}

		for j, v := range row {
			rowStore[v]++
			if rowStore[v] > 1 {
				return false
			}
			colStore[matrix[j][i]]++
			if colStore[matrix[j][i]] > 1 {
				return false
			}
		}
	}

	return true
}
