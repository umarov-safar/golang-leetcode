package main

import "strconv"

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		strNumber := strconv.Itoa(i)
		notSelfDivide := false

		for _, v := range strNumber {
			num, _ := strconv.Atoi(string(v))
			if num == 0 {
				notSelfDivide = true
				continue
			}
			if i%num != 0 {
				notSelfDivide = true
				break
			}
		}

		if !notSelfDivide {
			res = append(res, i)
		}
	}

	return res
}
