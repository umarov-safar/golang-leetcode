package phonecom

func letterCombinations(digits string) []string {
	digitsMap := getMappedDigits()
	result := []string{}

	var backtrack func(i int, letterCom string)

	backtrack = func(i int, letterCom string) {
		if len(letterCom) == len(digits) {
			result = append(result, letterCom)
			return
		}

		letters := digitsMap[string(digits[i])]

		for _, v := range letters {
			backtrack(i+1, letterCom+v)
		}
	}

	backtrack(0, "")

	return result
}

func getMappedDigits() map[string][]string {
	digitsMap := make(map[string][]string)
	digitsMap["2"] = []string{"a", "b", "c"}
	digitsMap["3"] = []string{"d", "e", "f"}
	digitsMap["4"] = []string{"g", "h", "i"}
	digitsMap["5"] = []string{"j", "k", "l"}
	digitsMap["6"] = []string{"m", "n", "o"}
	digitsMap["7"] = []string{"p", "q", "r", "s"}
	digitsMap["8"] = []string{"t", "u", "v"}
	digitsMap["9"] = []string{"w", "x", "y", "z"}

	return digitsMap
}
