package main

func scoreBalance(s string) bool {
	ps := make([]int, len(s))
	ps[0] = int(s[0]) - 96

	for i := 1; i < len(s); i++ {
		ps[i] = ps[i-1] + (int(s[i]) - 96)
	}

	for i := 0; i < len(ps); i++ {
		left := ps[i]
		right := ps[len(s)-1] - left

		if left == right {
			return true
		}
	}

	return false
}
