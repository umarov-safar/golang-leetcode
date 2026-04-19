package sumofprefixscoresofstrings

// Title: 2416. Sum of Prefix Scores of Strings
// Link: https://leetcode.com/problems/sum-of-prefix-scores-of-strings/

func sumPrefixScores(words []string) []int {
	root := &TrieNode{
		children: make(map[rune]*TrieNode),
	}
	for _, word := range words {
		insertWordToTrie(root, word)
	}

	res := make([]int, len(words))
	for i, word := range words {
		res[i] = countVisitedPrefix(root, word)
	}

	return res
}

type TrieNode struct {
	children map[rune]*TrieNode
	visited  int
}

func insertWordToTrie(node *TrieNode, word string) {
	for _, char := range word {

		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				visited:  0,
			}
		}

		node = node.children[char]
		node.visited++
	}
}

func countVisitedPrefix(node *TrieNode, word string) int {
	visited := 0
	for _, char := range word {
		node = node.children[char]
		visited += node.visited
	}

	return visited
}
