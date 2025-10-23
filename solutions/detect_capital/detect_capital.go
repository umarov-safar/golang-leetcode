package detectcapital

func detectCapitalUse(word string) bool {
    n := len(word)
    if n <= 1 {
        return true
    }

    if isCapital(word[0]) {
        if isCapital(word[1]) {
            return checkAllCapital(word)
        }
        return checkFirstCapital(word)
    }
    return checkAllLower(word)
}

func checkAllCapital(word string) bool {
    for i := 0; i < len(word); i++ {
        if !isCapital(word[i]) {
            return false
        }
    }
    return true
}

func checkFirstCapital(word string) bool {
    for i := 1; i < len(word); i++ {
        if isCapital(word[i]) {
            return false
        }
    }
    return true
}

func checkAllLower(word string) bool {
    for i := 0; i < len(word); i++ {
        if isCapital(word[i]) {
            return false
        }
    }
    return true
}

func isCapital(b byte) bool {
    return b >= 'A' && b <= 'Z'
}
