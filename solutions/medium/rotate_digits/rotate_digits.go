package rotatedigits

func rotatedDigits(n int) int {
	good := map[int]bool{
		2: true,
		5: true,
		6: true,
		9: true,
	}

	invalid := map[int]bool{
		3: true,
		4: true,
		7: true,
	}

	count := 0
	for i := 1; i <= n; i++ {
		digits := makeDigits(i)
		isValid := true
		hasDiff := false
		for _, num := range digits {
			if _, ok := invalid[num]; ok {
				isValid = false
			}

			if _, ok := good[num]; ok {
				hasDiff = true
			}
		}

		if isValid && hasDiff {
			count++
		}
	}

	return count
}

func makeDigits(x int) []int {
	if x <= 11 {
		return []int{x}
	}

	div := 10
	for div*10 <= x {
		div *= 10
	}

	digits := make([]int, 0)
	for div > 0 {
		num := x / div
		x = x % div
		div /= 10

		digits = append(digits, num)
	}

	return digits
}
