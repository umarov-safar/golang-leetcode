package design

import (
	"math/rand"
	"time"
)

type Solution struct {
	originalData   []int
	randomizedData []int
	r              *rand.Rand
}

func Constructor(nums []int) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	rData := make([]int, len(nums))
	copy(rData, nums)

	return Solution{
		originalData:   nums, 
		randomizedData: rData,
		r:              r,
	}
}

func (s *Solution) Reset() []int {
	return s.originalData
}

func (s *Solution) Shuffle() []int {
	for i := len(s.randomizedData) - 1; i > 0; i-- {
		j := s.r.Intn(i + 1)
		s.randomizedData[i], s.randomizedData[j] = s.randomizedData[j], s.randomizedData[i]
	}
	return s.randomizedData
}
