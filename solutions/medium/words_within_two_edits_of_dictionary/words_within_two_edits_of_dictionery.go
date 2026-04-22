package wordswithintwoeditsofdictionary

func twoEditWords(queries []string, dictionary []string) []string {
	res := make([]string, 0)
	for _, query := range queries {

		for _, dict := range dictionary {
			diff := 0
			for i := 0; i < len(query); i++ {
				if query[i] != dict[i] {
					diff++

					if diff > 2 {
						break
					}
				}
			}

			if diff <= 2 {
				res = append(res, query)
			}
		}

	}

	return res
}
