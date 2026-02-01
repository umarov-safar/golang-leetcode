package main

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefixSum := make([][]int, len(matrix))

	for i, raw := range matrix {
		ps := make([]int, len(raw))
		ps[0] = raw[0]

		for i := 1; i < len(raw); i++ {
			ps[i] = raw[i] + ps[i-1]
		}

		prefixSum[i] = ps
	}

	return NumMatrix{prefixSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for row2 >= row1 {
		if col1 == 0 {
			sum += this.matrix[row2][col2]
		} else {
			sum += this.matrix[row2][col2] - this.matrix[row2][col1-1]
		}
		row2--
	}

	return sum
}
