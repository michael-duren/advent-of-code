package lib

import (
	"fmt"
	"strconv"
)

func AsciToIntegerSlice(s []string) []int {
	intSlice := make([]int, 0, len(s))

	for _, l := range s {
		fmt.Println(l)
		result, err := strconv.Atoi(l)
		if err != nil {
			panic("unable to convert puzzle input")
		}
		intSlice = append(intSlice, result)
	}
	return intSlice

}
