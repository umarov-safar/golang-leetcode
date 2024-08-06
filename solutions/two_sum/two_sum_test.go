package two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	target := 9
	input1 := []int{2, 7, 11, 15}
	result := TwoSum(input1, target)

	fmt.Println(result)
	if result[0] != 0 || result[1] != 1 {
		t.Errorf("Not passed!")
	}

	target = 6
	input2 := []int{3, 2, 4}
	result = TwoSum(input2, target)
	if result[0] != 1 || result[1] != 2 {
		t.Errorf("Not passed!")
	}
}
