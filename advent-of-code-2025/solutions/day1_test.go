package solutions

import (
	"testing"
)

func Test_DayOnePartOne(t *testing.T) {
	tests := []struct {
		name string
		// function input
		want      int
		mainInput bool
	}{
		{
			name:      "correctly computes day one part one",
			want:      3,
			mainInput: false,
		},
		{
			name:      "correctly computes day one part one main input",
			want:      1141,
			mainInput: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DayOnePartOne(tt.mainInput)

			// Check result
			if got != tt.want {
				t.Errorf("DayOnePartOne() = %v, want %v", got, tt.want)
			}
		})

	}
}

func Test_DayOnePartTwo(t *testing.T) {
	tests := []struct {
		name string
		// function input
		want      int
		mainInput bool
	}{
		{
			name:      "correctly computes day one part two",
			want:      6,
			mainInput: false,
		},
		{
			name:      "correctly computes day one part two main input",
			want:      6634,
			mainInput: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DayOnePartTwo(tt.mainInput)

			// Check result
			if got != tt.want {
				t.Errorf("DayOnePartOne() = %v, want %v", got, tt.want)
			}
		})

	}
}
