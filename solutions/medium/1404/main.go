package main

func numSteps(s string) int {
	steps := 0

	for len(s) > 1 {
		if string(s[len(s)-1]) == "0" {
			s = s[:len(s)-1]
		} else {
			s = addOne(s)
		}
		steps++
	}

	return steps
}

func addOne(s string) string {
	i := len(s) - 1

	for i >= 0 && string(s[i]) != "0" {
		s = s[:i] + "0" + s[i+1:]
		i--
	}

	if i < 0 {
		s = "1" + s
	} else {
		s = s[:i] + "1" + s[i+1:]
	}

	return s
}
