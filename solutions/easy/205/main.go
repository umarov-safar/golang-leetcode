package main

func isIsomorphic(s string, t string) bool {
	smap := make(map[byte]byte)
	tmap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		tchar := t[i]
		schar := s[i]

		sv, oks := smap[s[i]]

		if oks && (sv != tchar) {
			return false
		}

		ttchar, ok := tmap[tchar]

		if ok && ttchar != schar {
			return false
		}

		smap[s[i]] = tchar
		tmap[t[i]] = schar
	}

	return true
}
