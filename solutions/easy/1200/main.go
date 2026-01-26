package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minAbs := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		minAbs = min(minAbs, arr[i+1]-arr[i])
	}

	ans := make([][]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		if minAbs == arr[i+1]-arr[i] {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}

	return ans
}
