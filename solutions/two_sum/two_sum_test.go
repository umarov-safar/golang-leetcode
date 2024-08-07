package two_sum

import (
	"fmt"
	"testing"
)

var inputs = [][]int{
	{2, 7, 11, 15},
	{3, 2, 4},
	{3, 3},
}
var targets = []int{9, 6, 6}

func TestTwoSum(t *testing.T) {
	for iInput, input := range inputs {
		target := targets[iInput]
		res := TwoSum(input, target)
		sum := input[res[0]] + input[res[1]]
		if sum != target {
			t.Errorf("Expected %d given %d", target, sum)
		}
	}

	fmt.Println(inputs, targets)
}

func TestTwoSumWithHash(t *testing.T) {
	for iInput, input := range inputs {
		target := targets[iInput]
		res := TwoSumWithHash(input, target)
		sum := input[res[0]] + input[res[1]]
		if sum != target {
			t.Errorf("Expected %d given %d", target, sum)
		}
	}
}
