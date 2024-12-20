package d04

import (
	"awesomeProject/util"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 4,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "18", Real: "2569"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "9", Real: "1998"},
		},
	}
}

func Part1(input string, example bool) string {
	grid := parse(input)

	count := 0
	for y, row := range grid {
		for x := range row {
			count += findWord("XMAS", grid, util.Point{x, y}, util.Point{x, y})
		}
	}

	return strconv.Itoa(count)
}

func Part2(input string, example bool) string {
	grid := parse(input)

	count := 0
	for y, row := range grid {
		if y == 0 || y == len(grid)-1 {
			continue
		}
		for x := range row {
			if x == 0 || x == len(row)-1 {
				continue
			}
			if row[x] == "A" {
				tl, tr, bl, br := grid[y-1][x-1], grid[y-1][x+1], grid[y+1][x-1], grid[y+1][x+1]
				firstDiagonal := (tl == "M" && br == "S") || (tl == "S" && br == "M")
				secondDiagonal := (tr == "M" && bl == "S") || (tr == "S" && bl == "M")
				if firstDiagonal && secondDiagonal {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count)
}

var deltas = []util.Point{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func findWord(word string, grid [][]string, pos util.Point, prev util.Point) (count int) {
	if pos.X < 0 || pos.Y < 0 || pos.Y >= len(grid) || pos.X >= len(grid[pos.Y]) {
		return 0
	}
	if grid[pos.Y][pos.X] != string(word[0]) {
		return 0
	}
	if len(word) == 1 {
		return 1
	}
	if pos == prev {
		for _, delta := range deltas {
			count += findWord(word[1:], grid, util.Point{pos.X + delta.X, pos.Y + delta.Y}, pos)
		}
	} else {
		dx, dy := pos.X-prev.X, pos.Y-prev.Y
		count += findWord(word[1:], grid, util.Point{pos.X + dx, pos.Y + dy}, pos)
	}
	return
}

func parse(str string) (rows [][]string) {
	for _, line := range strings.Split(str, "\n") {
		rows = append(rows, strings.Split(line, ""))
	}
	return
}
