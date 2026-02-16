package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	n := len(potions)

	for i := range spells {
		findFirstGreat := binarySearch(potions, spells[i], success)

		if findFirstGreat < 0 {
			spells[i] = 0
		} else {
			spells[i] = n - findFirstGreat
		}
	}

	return spells
}

func binarySearch(potions []int, spell int, success int64) int {
	l, r := 0, len(potions)-1
	for l <= r {
		mid := l + ((r - l) / 2)
		product := int64(potions[mid] * spell)

		if product < success && mid < r && int64(potions[mid+1]*spell) > success {
			return mid + 1
		} else if product >= success && mid > 0 && int64(potions[mid-1]*spell) < success {
			return mid
		} else if mid == 0 && product >= success {
			return mid
		} else if product >= success {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
