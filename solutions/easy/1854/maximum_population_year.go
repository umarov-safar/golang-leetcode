package maximumpopulationyear

func maximumPopulation(logs [][]int) int {
	years := make(map[int]int)

	for _, row := range logs {
		for i := row[0]; i < row[1]; i++ {
			years[i]++
		}
	}

	earlyYear := 0
	mxCount := 0
	for year, count := range years {
		if count > mxCount {
			earlyYear = year
			mxCount = count	
		} else if count == mxCount && year < earlyYear {
			earlyYear = year
		}
	}

	return earlyYear
}
