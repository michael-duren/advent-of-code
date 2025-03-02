package main

import (
	"advent-of-code-2024/day1"
	"advent-of-code-2024/day2"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
)

func main() {
	day := os.Args[1]
	dayParsed, err := strconv.Atoi(day)
	if err != nil {
		log.Fatalf("The argument: %s is not a valid day. error: %v", day, err)
	}

	switch dayParsed {
	case 1:
		day1.SolvePartOne()
	case 2:
		day2.Solve()
	default:
		log.Fatalf("the number was too high or too low. no day corresponds the number input: %d", dayParsed)
	}
}
