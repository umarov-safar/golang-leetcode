package sortnumbersbyitsbits

import "sort"

func sortByBits(arr []int) []int {
	bitCount := make(map[int]int)

	count := func(n int) int {
		if v, ok := bitCount[n]; ok {
			return v
		}
		c := 0
		x := n
		for x != 0 {
			c += x & 1
			x >>= 1
		}
		bitCount[n] = c
		return c
	}

	sort.SliceStable(arr, func(i, j int) bool {
		bi := count(arr[i])
		bj := count(arr[j])

		if bi == bj {
			return arr[i] < arr[j]
		}
		return bi < bj
	})

	return arr
}
