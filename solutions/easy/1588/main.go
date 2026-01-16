package main

func sumOddLengthSubarrays(arr []int) int {
	arr = sumPrefix(arr)
	s := arr[len(arr)-1]
	for i := 2; i < len(arr); i++ {
		if i%2 != 0 {
			continue
		}
		l, r := 0, i
		for r < len(arr) && (r-1)-l <= i {
			if l == 0 {
				s += arr[i]
				r++
				l++
				continue
			}
			s += arr[r] - arr[l-1]
			l++
			r++
		}
	}

	return s
}

func sumPrefix(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i]
	}

	return arr
}
