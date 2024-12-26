package d25

import (
	"awesomeProject/util"
	"fmt"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 25,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "3", Real: "3196"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "???", Real: "???"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)

	matches := 0
	for _, lock := range data.locks {
		for _, key := range data.keys {
			if match(lock, key) {
				matches++
			}
		}
	}

	return strconv.Itoa(matches)
}

func Part2(input string, example bool) string {
	data := parse(input)
	fmt.Println(data)
	return "TODO"
}

func match(lock, key []int) bool {
	if len(lock) != len(key) {
		return false
	}
	for x := range len(lock) {
		if lock[x]+key[x] >= 6 {
			return false
		}
	}
	return true
}

func parse(input string) (data Data) {
	parts := strings.Split(input, "\n\n")
	for _, part := range parts {
		arr := util.ParseCharMatrix(part)

		cols := make([]int, len(arr[0]))
		for x := range len(arr[0]) {
			cols[x] = -1 // Ignore top/bottom row
			for y := range len(arr) {
				if arr[y][x] == "#" {
					cols[x]++
				}
			}
		}

		if arr[0][0] == "#" {
			data.locks = append(data.locks, cols)
		} else {
			data.keys = append(data.keys, cols)
		}
	}

	return
}

type Data struct {
	locks [][]int
	keys  [][]int
}
