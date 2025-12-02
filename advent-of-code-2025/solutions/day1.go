package solutions

import (
	"2025/lib"
	"strconv"
)

func DayOnePartOne(realInput bool) int {
	input := lib.GetPuzzleInput(1, realInput)
	dial := 50
	numberOfZeros := 0

	for _, line := range input {
		direction := line[0]
		amount := line[1:]
		intAmount, err := strconv.Atoi(amount)
		if err != nil {
			panic("unable to convert amount")
		}
		if direction == 'L' {
			intAmount *= -1
		}

		dial = (dial + intAmount) % 100
		if dial == 0 {
			numberOfZeros++
		}
	}

	return numberOfZeros
}

func DayOnePartTwo(realInput bool) int {
	input := lib.GetPuzzleInput(1, realInput)
	dial := 50
	numberOfZeros := 0

	for _, line := range input {
		direction := line[0]
		amount := line[1:]
		intAmount, err := strconv.Atoi(amount)
		if err != nil {
			panic("unable to convert amount")
		}
		if direction == 'R' {
			numberOfZeros += (dial + intAmount) / 100
			dial = (dial + intAmount) % 100
		} else { // 'L'
			if dial == 0 {
				numberOfZeros += intAmount / 100
			} else if intAmount >= dial {
				numberOfZeros += 1 + (intAmount-dial)/100
			}
			dial = ((dial-intAmount)%100 + 100) % 100
		}
	}

	return numberOfZeros
}
