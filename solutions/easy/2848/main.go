package main

func numberOfPoints(nums [][]int) int {
	ar := make([]int, 102)
	for _, val := range nums {
		ar[val[0]]++
		ar[val[1]+1]--
	}

	ans := 0
	curr := 0
	for i := range ar {
		curr += ar[i]
		if curr > 0 {
			ans++
		}
	}
	return ans
}
