package countmajoritysubarray

func countMajoritySubarrays(nums []int, target int) int {
	var majority int
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == target {
				count++
			}

			if count*2 > j-i+1 {
				majority++
			}
		}
	}

	return majority
}
