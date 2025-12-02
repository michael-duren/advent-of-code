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

// package solutions
//
// import (
// 	"2025/lib"
// 	"strconv"
// 	"strings"
// )
//
// type IDRange struct {
// 	start int64
// 	end   int64
// }
//
// func DayTwoPartOne(realInput bool) int64 {
// 	input := lib.GetPuzzleInput(2, realInput)
// 	if len(input) != 1 {
// 		panic("unexpected input length")
// 	}
// 	idRanges := strings.Split(input[0], ",")
// 	ranges := make([]IDRange, 0, len(idRanges))
// 	for _, idRange := range idRanges {
// 		ranges = append(ranges, parseIDRange(idRange))
// 	}
// 	invalidIds := make(map[int64]bool)
// 	for _, r := range ranges {
// 		for id := r.start; id <= r.end; id++ {
// 			if !isValidID(id) {
// 				invalidIds[id] = true
// 			}
// 		}
// 	}
//
// 	var total int64
// 	for id := range invalidIds {
// 		total += id
// 	}
// 	return total
// }
//
//
// func parseIDRange(idRange string) IDRange {
// 	parts := strings.Split(idRange, "-")
// 	if len(parts) != 2 {
// 		panic("unexpected range format")
// 	}
// 	start, err := strconv.ParseInt(parts[0], 10, 64)
// 	if err != nil {
// 		panic("unable to convert range start")
// 	}
// 	end, err := strconv.ParseInt(parts[1], 10, 64)
// 	if err != nil {
// 		panic("unable to convert range end")
// 	}
// 	return IDRange{start: start, end: end}
//
// }
