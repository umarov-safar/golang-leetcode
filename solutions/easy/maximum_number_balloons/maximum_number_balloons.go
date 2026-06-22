func maxNumberOfBalloons(text string) int {
	m := make(map[rune]int)

	for _, char := range text {
		if char == 'b' {
			m[char]++
		}

		if char == 'a' {
			m[char]++
		}

		if char == 'l' {
			m[char]++
		}

		if char == 'o' {
			m[char]++
		}

		if char == 'n' {
			m[char]++
		}
	}

	count := 0
	for {
		if v, ok := m['b']; ok && v > 0 {
			m['b']--
		} else {
			return count
		}

		if v, ok := m['a']; ok && v > 0 {
			m['a']--
		} else {
			return count
		}

		if v, ok := m['l']; ok && v > 1 {
			m['l'] -= 2
		} else {
			return count
		}

		if v, ok := m['o']; ok && v > 1 {
			m['o'] -= 2
		} else {
			return count
		}

		if v, ok := m['n']; ok && v > 0 {
			m['n']--
		} else {
			return count
		}

		count++
	}
}


func maxNumberOfBalloonsOptimized(text string) int {
    m := make(map[rune]int)

    for _, char := range text {
        if char == 'b' {
            m[char]++
        }
        
        if char == 'a' {
            m[char]++
        }

        if char == 'l' {
            m[char]++
        }

        if char == 'o' {
            m[char]++
        }

        if char == 'n' {
            m[char]++
        }
     }

    if len(m) < 5 {
        return 0 
    }
    
    minimum := math.MaxInt

    for char, v  := range m {
        if char == 'l' || char == 'o' {
            v = v / 2
        } 

        if minimum > v {
            minimum = v
        }
    }

    return minimum
}