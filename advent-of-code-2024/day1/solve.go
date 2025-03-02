package day1

import (
	"advent-of-code-2024/util"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type lists struct {
	a []int
	b []int
}

func parseCol(str string) (int, error) {
	strTrimmed := strings.ReplaceAll(str, " ", "")
	val, err := strconv.Atoi(strTrimmed)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func splitRow(row string) ([]string, error) {
	cols := strings.Split(row, " ")
	if len(cols) < 2 {
		return nil, fmt.Errorf("cols: %v does not have the right amount of return vals. expected 2 or more got %d", cols, len(cols))
	}
	parsedCols := make([]string, 0, 2)
	for _, v := range cols {
		parsed := strings.ReplaceAll(v, " ", "")

		if parsed == "" {
			continue
		}
		parsedCols = append(parsedCols, parsed)
	}
	return parsedCols, nil
}

func newLists(str string) *lists {
	rows := strings.Split(str, "\n")
	size := len(rows)/2 + 1
	a := make([]int, 0, size)
	b := make([]int, 0, size)
	for _, row := range rows {
		if row == "" {
			continue
		}
		cols, err := splitRow(row)
		if err != nil {
			log.Fatalf("error splitting cols: %v", err)
		}
		aVal, err := parseCol(cols[0])
		if err != nil {
			log.Fatalf("error parsing a val: %v", err)
		}
		bVal, err := parseCol(cols[1])
		if err != nil {
			log.Fatalf("error parsing b val %v", err)
		}
		a = append(a, aVal)
		b = append(b, bVal)
	}

	sort.Ints(a)
	sort.Ints(b)
	return &lists{a, b}
}

func (l *lists) sumDistances() (int, error) {
	if len(l.a) != len(l.b) {
		return 0, fmt.Errorf("lengths of a and b are not the same. length of a: %d length of b: %d", l.a, l.b)
	}
	sum := 0
	for i, v := range l.a {
		sum = sum + int(math.Abs(float64(v-l.b[i])))
	}
	return sum, nil
}

func (l *lists) sumByRepatedNumbers() (int, error) {
	if len(l.a) != len(l.b) {
		return 0, fmt.Errorf("lengths of a and b are not the same. length of a: %d length of b: %d", l.a, l.b)
	}

	sum := 0
	for _, v := range l.a {
		repeated := 0
		for _, vv := range l.b {
			if vv == v {
				repeated++
			}
		}
		sum = sum + (repeated * v)
	}

	return sum, nil
}

func SolvePartOne() {
	log.Info("Running day one part one")
	cwd, _ := os.Getwd()
	log.Infof("Current directory: %s", cwd)
	puzzleStr := util.ReadFromFile("day1/puzzle-input.txt")
	l := newLists(puzzleStr)
	sum, err := l.sumDistances()
	if err != nil {
		log.Fatalf("there was an error summing the lists: %v", err)
	}
	diffSummed, err := l.sumByRepatedNumbers()
	if err != nil {
		log.Fatalf("there was an error summing the diffed lists: %v", err)
	}
	log.Infof("successfully summed lists: %d", sum)
	log.Infof("successfully summed diffed lists: %d", diffSummed)
}
