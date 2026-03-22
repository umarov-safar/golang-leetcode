package main

func firstUniqChar(s string) int {

	m := make(map[rune]int)
	i := 0
	for _, val := range s {
		m[val]++

		count, _ := m[rune(s[i])]
		if count > 1 {
			i++
		}
	}

	c, _ := m[rune(s[i])]
	for c > 1 && i < len(s) {
		i++
		if i == len(s) {
			return -1
		}
		c, _ = m[rune(s[i])]
	}

	return i
}
