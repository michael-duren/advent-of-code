package day2

import (
	"advent-of-code-2024/util"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type matrix struct {
	levels [][]int
}

func newMatrix(lines []string) *matrix {
	m := make([][]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		strSplit := strings.Split(line, " ")
		nums := make([]int, 0, len(strSplit))
		for _, v := range strSplit {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("error converting ascii to int: %v", err)
			}
			nums = append(nums, n)
		}
		m = append(m, nums)
	}
	return &matrix{levels: m}
}

func (m *matrix) isLineSafe(i int) (bool, int) {
	line := m.levels[i]
	if len(line) < 2 {
		return true, 0
	}

	desc := line[0] > line[1]
	asc := !desc

	if math.Abs(float64(line[0]-line[1])) > 3 {
		return false, 1
	}

	for i, item := range line {
		// there is no nextitem
		if i+1 == len(line) {
			return true, i + 1
		}
		nextItem := line[i+1]
		diff := math.Abs(float64(item - nextItem))
		if diff > 3 {
			return false, i + 1
		}
		if desc && item < nextItem {
			return false, i + 1
		}

		if asc && item > nextItem {
			return false, i + 1
		}

		if item == nextItem {
			return false, i + 1
		}
	}
	return true, i + 1
}

func tryRemoveLevel(row *[]int, level int) bool {
	if level < 0 || level >= len(*row) {
		return false
	}
	*row = slices.Delete((*row), level, level+1)
	return true
}

func (m *matrix) countSafe() int {
	var total int
	for i := range m.levels {
		safe, _ := m.isLineSafe(i)
		if safe {
			total += 1
		}
	}
	return total
}

func (m *matrix) countSafeWithDampener() int {
	var total int
	for i := range m.levels {
		safe, l := m.isLineSafe(i)
		if !safe {
			tryRemoveLevel(&m.levels[i], l)
			safe, _ = m.isLineSafe(i)
		}
		log.Infof("row: %v, result: %t", &m.levels[i], safe)
		if safe {
			total++
		}
	}
	return total
}

func Solve() {
	log.Info("solving day 2")
	lines := util.ReadFileToLines("day2/puzzle-input.txt")
	// lines := util.ReadFileToLines("day2/example-input.txt")
	m := newMatrix(lines)
	totalSafe := m.countSafe()
	totalDamperedSafe := m.countSafeWithDampener()
	log.Infof("Total safe for part one: %d", totalSafe)
	log.Infof("Total dampener safe for part two: %d", totalDamperedSafe)
}
