package easy

// Task: 796. Rotate String
// Link: https://leetcode.com/problems/rotate-string/description

func rotateString(s string, goal string) bool {
	if s == goal {
		return true
	}

	t := s

	for i := range s {
		char := string(s[i])

		t = t[1:] + char
		if t == goal {
			return true
		}
	}

	return false
}
