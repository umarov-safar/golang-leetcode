package countvalidwordoccurrences

import "strings"

// Titile: 3926. Count Valid Word Occurrences
// Link: https://leetcode.com/problems/count-valid-word-occurrences/description/

func countWordOccurrences(chunks []string, queries []string) []int {
	words := map[string]int{}
	text := strings.Join(chunks, "")

	newText := strings.ReplaceAll(text, "--", " ")

	chunks = strings.Fields(newText)

	for _, word := range chunks {
		if word == "-" {
			continue
		}
		if word[0] == '-' && word[len(word)-1] == '-' {
			word = word[1 : len(word)-1]
		} else if word[0] == '-' {
			word = word[1:]
		} else if word[len(word)-1] == '-' {
			word = word[:len(word)-1]
		}
		words[word]++
	}

	res := make([]int, len(queries))

	for i, q := range queries {
		if v, ok := words[q]; ok {
			res[i] = v
		} else {
			res[i] = 0
		}
	}

	return res
}
