package easy

import "testing"

func TestCheckValid(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   bool
	}{
		{
			name: "Valid 3x3 matrix",
			matrix: [][]int{
				{1, 2, 3},
				{3, 1, 2},
				{2, 3, 1},
			},
			want: true,
		},
		{
			name: "Invalid row duplicate",
			matrix: [][]int{
				{1, 2, 2},
				{3, 1, 2},
				{2, 3, 1},
			},
			want: false,
		},
		{
			name: "Invalid column duplicate",
			matrix: [][]int{
				{1, 2, 3},
				{3, 1, 2},
				{1, 3, 1},
			},
			want: false,
		},
		{
			name: "Valid 2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{2, 1},
			},
			want: true,
		},
		{
			name: "Invalid 2x2 matrix",
			matrix: [][]int{
				{1, 1},
				{2, 2},
			},
			want: false,
		},
		{
			name: "Single element matrix",
			matrix: [][]int{
				{1},
			},
			want: true,
		},
		{
			name:   "Empty matrix",
			matrix: [][]int{},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckValid(tt.matrix)
			if got != tt.want {
				t.Errorf("CheckValid() = %v, want %v", got, tt.want)
			}
			got2 := CheckValidV2(tt.matrix)
			if got2 != tt.want {
				t.Errorf("CheckValidV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
