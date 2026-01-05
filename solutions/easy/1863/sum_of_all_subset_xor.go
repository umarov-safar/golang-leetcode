package sumsubsetxor

func subsetXORSum(nums []int) int {
	s := 0
	var sumXOR func(i int, path []int)

	sumXOR = func(i int, path []int) {
		if i == len(nums) {
			if len(path) == 0 {
				return
			} else if len(path) == 1 {
				s += path[0]
				return
			}
			x := path[0]
			for j := 1; j < len(path); j++ {
				x = (x ^ path[j])
			}
			s += x
			return
		}

		path = append(path, nums[i])
		sumXOR(i+1, path)
		path = path[:len(path)-1]
		sumXOR(i+1, path)
	}

	sumXOR(0, []int{})

	return s

}
