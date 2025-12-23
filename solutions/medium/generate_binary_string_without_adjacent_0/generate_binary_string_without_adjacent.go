package generatebinarystringwithoutadjacent0

func ValidStrings(n int) []string {
	var res []string
	backtrack("", n, &res)
	return res
}

func backtrack(s string, n int, res *[]string) {
	if len(s) == n {
		*res = append(*res, s)
		return
	}

	if len(s) == 0 {
		backtrack(s+"0", n, res)
	} else if string(s[len(s)-1]) != "0" {
		backtrack(s+"0", n, res)
	}

	backtrack(s+"1", n, res)
}
