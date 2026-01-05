package fourdivisors

func sumFourDivisors(nums []int) int {
	sum := 0
	for _, v := range nums {
		if v <= 4 {
			continue
		}
		s := 1 + v
		j := 2
		for i := 2; i <= int(v/2)+1; i++ {
			if v%i == 0 {
				s += i
				j++
			}
			if j > 4 {
				break
			}
		}

		if j == 4 {
			sum += s
		}
	}

	return sum
}
