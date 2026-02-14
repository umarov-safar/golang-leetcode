package main

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	jw := make(map[rune]struct{})
	for _, char := range jewels {
		jw[char] = struct{}{}
	}

	for _, v := range stones {
		if _, ok := jw[v]; ok {
			count++
		}
	}

	return count
}