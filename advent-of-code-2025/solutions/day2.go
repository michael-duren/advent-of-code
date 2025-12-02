package solutions

import (
	"2025/lib"
	"strconv"
	"strings"
)

type IDRange struct {
	start int64
	end   int64
}

func DayTwoPartOne(realInput bool) int64 {
	input := lib.GetPuzzleInput(2, realInput)
	if len(input) != 1 {
		panic("unexpected input length")
	}
	idRanges := strings.Split(input[0], ",")
	ranges := make([]IDRange, 0, len(idRanges))
	for _, idRange := range idRanges {
		ranges = append(ranges, parseIDRange(idRange))
	}
	invalidIds := make(map[int64]bool)
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			if !isValidID(id) {
				invalidIds[id] = true
			}
		}
	}

	var total int64
	for id := range invalidIds {
		total += id
	}
	return total
}

func DayTwoPartTwo(realInput bool) int64 {
	input := lib.GetPuzzleInput(2, realInput)
	if len(input) != 1 {
		panic("unexpected input length")
	}
	idRanges := strings.Split(input[0], ",")
	ranges := make([]IDRange, 0, len(idRanges))
	for _, idRange := range idRanges {
		ranges = append(ranges, parseIDRange(idRange))
	}
	invalidIds := make(map[int64]bool)
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			if !isValidIDTwo(id) {
				invalidIds[id] = true
			}
		}
	}

	var total int64
	for id := range invalidIds {
		total += id
	}
	return total
}

func isValidID(id int64) bool {
	idStr := strconv.FormatInt(id, 10)
	if len(idStr)%2 != 0 {
		return true // odd length can never be X+X
	}
	half := len(idStr) / 2
	return idStr[:half] != idStr[half:]
}

func isValidIDTwo(id int64) bool {
	// 1188511885 false
	// 11 or 22 false
	idStr := strconv.FormatInt(id, 10)
	// we want to split and compare essentially dividing by 2, 3, 4, up to the length
	for split := 1; split <= len(idStr)/2; split++ {
		splits := []string{}

		for i := 0; i < len(idStr); i += split {
			end := min(i+split, len(idStr))
			splits = append(splits, idStr[i:end])
		}
		allSame := true
		first := splits[0]
		for _, s := range splits[1:] {
			if s != first {
				allSame = false
				break
			}
		}
		// if all same then the id is invalid
		if allSame {
			return false
		}
	}
	return true
}

func parseIDRange(idRange string) IDRange {
	parts := strings.Split(idRange, "-")
	if len(parts) != 2 {
		panic("unexpected range format")
	}
	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic("unable to convert range start")
	}
	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic("unable to convert range end")
	}
	return IDRange{start: start, end: end}

}
