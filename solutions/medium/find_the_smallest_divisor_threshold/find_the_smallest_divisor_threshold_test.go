package midlle

import (
	"testing"
)

type Input struct {
	nums      []int
	threshold int
	expected  int
}

func TestSmallestDivisor(t *testing.T) {
	data := []Input{
		{[]int{1, 2, 5, 9}, 6, 5},
		{[]int{44, 22, 33, 11, 1}, 5, 44},
	}

	for _, input := range data {
		result := SmallestDivisor(input.nums, input.threshold)
		if result != input.expected {
			t.Errorf("Expect %v given %v", input.expected, result)
		}
	}
}
