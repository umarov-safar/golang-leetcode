package design

import (
	"reflect"
	"testing"
)

func TestSolutionMemoryEfficient(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	s := Constructor(original)

	reset := s.Reset()
	if !reflect.DeepEqual(reset, original) {
		t.Errorf("Reset() = %v; want %v", reset, original)
	}

	shuffled := s.Shuffle()
	if reflect.DeepEqual(shuffled, original) {
		t.Errorf("Shuffle() returned the same as original: %v", shuffled)
	}

	reset2 := s.Reset()
	if !reflect.DeepEqual(reset2, original) {
		t.Errorf("Reset() after shuffle = %v; want %v", reset2, original)
	}

	shuffle1 := s.Shuffle()
	shuffle2 := s.Shuffle()
	if reflect.DeepEqual(shuffle1, shuffle2) {
		t.Logf("Warning: two consecutive shuffles were identical (possible but unlikely): %v", shuffle1)
	}

	if _, ok := any(s).(Solution); !ok {
		t.Errorf("Constructor did not return Solution type")
	}
}
