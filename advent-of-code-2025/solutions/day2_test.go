package solutions

import (
	"testing"
)

func Test_DayTwoPartOne(t *testing.T) {
	tests := []struct {
		name string
		// function input
		want      int64
		mainInput bool
	}{
		{
			name:      "correctly computes day two part one",
			want:      1227775554,
			mainInput: false,
		},
		{
			name:      "correctly computes day one part one main input",
			want:      19219508902,
			mainInput: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DayTwoPartOne(tt.mainInput)

			// Check result
			if got != tt.want {
				t.Errorf("DayOnePartOne() = %v, want %v", got, tt.want)
			}
		})

	}
}

func Test_DayTwoPartTwo(t *testing.T) {
	tests := []struct {
		name string
		// function input
		want      int64
		mainInput bool
	}{
		{
			name:      "correctly computes day two part one",
			want:      4174379265,
			mainInput: false,
		},
		{
			name:      "correctly computes day one part one main input",
			want:      19219508902,
			mainInput: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DayTwoPartTwo(tt.mainInput)

			// Check result
			if got != tt.want {
				t.Errorf("DayOnePartOne() = %v, want %v", got, tt.want)
			}
		})

	}
}
