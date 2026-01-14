package main

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	tm := make(map[byte]int)

	for i := range t {
		tm[t[i]]++
	}

	subMap := make(map[byte]int)
	res := s + "a" // greater than the s input to check even s has substring
	left := 0

	for right := range s {
		subMap[s[right]]++

		for right >= len(t)-1 && checkHasCharacters(subMap, tm) {
			if subMap[s[left]] > 0 {
				subMap[s[left]]--
			} else {
				delete(subMap, s[left])
			}
			t := s[left : right+1]
			if len(t) < len(res) {
				res = t
			}
			left++
		}
	}

	if len(res) > len(s) {
		return ""
	}

	return res
}

func checkHasCharacters(subMap map[byte]int, tm map[byte]int) bool {
	for char, count := range tm {
		if v, ok := subMap[char]; ok && v >= count {
			continue
		} else {
			return false
		}
	}

	return true
}
