package generatebinarystringwithoutadjacent0

import (
	"testing"
)

func TestValidStrings(t *testing.T) {
	tests := []struct {
		n      int
		expect []string
	}{
		{2, []string{"01", "10", "11"}},
		{3, []string{"101", "110", "011", "111", "010"}},
	}

	for _, tt := range tests {
		result := ValidStrings(tt.n)
		if len(result) != len(tt.expect) {
			t.Errorf("n=%d: expected %d results, got %d", tt.n, len(tt.expect), len(result))
		}
	}
}
