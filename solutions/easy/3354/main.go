package main

func countValidSelections(nums []int) int {
	s := sum(nums)

	left := 0
	ans := 0
	for _, v := range nums {
		if v == 0 && s-left == left {
			ans += 2
		} else if v == 0 && s-left == left-1 {
			ans += 1
		} else if v == 0 && s-1-left == left {
			ans += 1
		}

		left += v
	}

	return ans
}

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}

	return s
}
