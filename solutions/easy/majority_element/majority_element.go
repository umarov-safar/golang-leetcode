package majorityelement

func majorityElement(nums []int) int {
	res, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			res = num
		}

		if res == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return res
}
