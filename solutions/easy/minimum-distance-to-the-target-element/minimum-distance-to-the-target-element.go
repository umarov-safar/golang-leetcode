package minimumdistancetothetargetelement

func getMinDistance(nums []int, target int, start int) int {
	minDistance := len(nums)
	for i, n := range nums {
		if n == target {
			minDistance = min(abs(i-start), minDistance)
		}
	}

	return minDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
