package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	// m := len(queries)
	n := len(nums)
	ans := []int{}
	sort.Ints(nums)

	ps := make([]int, n)

	k := 1
	ps[0] = nums[0]
	for k < n {
		ps[k] = ps[k-1] + nums[k]
		k++
	}

	for _, val := range queries {
		for i, v := range ps {
			if v == val {
				ans = append(ans, i+1)
				break
			} else if v > val {
				ans = append(ans, i)
				break
			}

			if i == len(nums)-1 && val > ps[i] {
				ans = append(ans, i+1)
			}
		}

	}

	return ans
}
