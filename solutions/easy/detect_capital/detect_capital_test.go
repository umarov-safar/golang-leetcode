package detectcapital

import (
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {
	inputs := []string{"USA", "leetcode", "Google"}
	wrongInputs := []string{"USa", "lEetcode", "Gooogle"}

	for _, word := range inputs {
		if detectCapitalUse(word) != true {
			t.Error("Correct word detected as wrong")
		}
	}

	for _, word := range wrongInputs {
		if detectCapitalUse(word) == true {
			t.Error("Wrong word detected as correct")
		}
	}
}
