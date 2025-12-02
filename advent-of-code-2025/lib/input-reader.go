package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func readInput(day string) []string {
	day = fmt.Sprintf("%s.txt", day)
	path := filepath.Join("..", "test-input", day)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		panic("error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
		panic("error scanning file")
	}
	return lines
}

func GetPuzzleInput(day int, realInput bool) []string {
	var dayInput string
	if realInput {
		dayInput = fmt.Sprintf("day%d", day)
	} else {
		dayInput = fmt.Sprintf("day%d-s", day)
	}
	return readInput(dayInput)
}
