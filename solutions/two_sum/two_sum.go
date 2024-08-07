package two_sum

// With loop!
func TwoSum(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// With hash
func TwoSumWithHash(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		
		if index, ok := hash[ target - num]; ok {
			return []int{i, index}
		}
		hash[nums[i]] = i
	}

	return []int{}
}
