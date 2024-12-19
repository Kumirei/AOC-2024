package d10

import (
	"awesomeProject/util"
	"strconv"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 10,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "36", Real: "733"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "81", Real: "1514"},
		},
	}
}

var deltas = []util.Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func Part1(input string, example bool) string {
	data := util.Parse2DIntArray(input, "")
	width, height := len(data[0]), len(data)

	heads := findStarts(data)
	score := 0

	for _, head := range heads {
		queue := []util.Point{head}
		visited := util.NewSet[util.Point]()
		for len(queue) > 0 {
			var curr util.Point
			queue, curr = util.Pop(queue)
			if visited.Has(curr) {
				continue
			}
			visited.Add(curr)
			val := data[curr.Y][curr.X]
			for _, delta := range deltas {
				x, y := curr.X+delta.X, curr.Y+delta.Y
				if util.InBounds(width, height, x, y) && data[y][x] == val+1 {
					queue = append(queue, util.Point{x, y})
				}
			}
		}

		for p := range visited.List {
			if data[p.Y][p.X] == 9 {
				score++
			}
		}
	}

	return strconv.Itoa(score)
}

func Part2(input string, example bool) string {
	data := util.Parse2DIntArray(input, "")

	heads := findStarts(data)
	count := 0
	for _, head := range heads {
		count += countTrails(data, head)
	}

	return strconv.Itoa(count)
}

func countTrails(data [][]int, pos util.Point) (count int) {
	width, height := len(data[0]), len(data)
	val := data[pos.Y][pos.X]
	if data[pos.Y][pos.X] == 9 {
		return 1
	}
	for _, delta := range deltas {
		x, y := pos.X+delta.X, pos.Y+delta.Y
		if util.InBounds(width, height, x, y) && data[y][x] == val+1 {
			count += countTrails(data, util.Point{x, y})
		}
	}
	return
}

func findStarts(data [][]int) (starts []util.Point) {
	for y, row := range data {
		for x, cell := range row {
			if cell == 0 {
				starts = append(starts, util.Point{x, y})
			}
		}
	}
	return
}
