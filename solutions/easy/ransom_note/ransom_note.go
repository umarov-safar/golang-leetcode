package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for _, char := range magazine {
		m[char]++
	}

	for _, char := range ransomNote {
		val, ok := m[char]
		if ok == false || val == 0 {
			return false
		}

		m[char]--
	}

	return true
}
