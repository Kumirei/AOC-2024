package d18

import (
	"awesomeProject/util"
	"fmt"
	"strconv"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 18,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "22", Real: "280"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "6,1", Real: "28,56"},
		},
	}
}

func Part1(input string, example bool) string {
	data := util.Parse2DIntArray(input, ",")
	data = getSlice(data, example)
	width, height := getSize(example)
	m := coordsToMap(data, width, height)

	steps, _ := bfs(m)
	return strconv.Itoa(steps)
}

func Part2(input string, example bool) string {
	data := util.Parse2DIntArray(input, ",")
	width, height := getSize(example)

	for i := range len(data) {
		slice := data[:i+1]
		m := coordsToMap(slice, width, height)
		_, ok := bfs(m)
		if !ok {
			return fmt.Sprintf("%d,%d", data[i][0], data[i][1])
		}
	}

	panic("Unreachable")
}

func bfs(m [][]string) (steps int, ok bool) {
	width, height := len(m[0]), len(m)
	start := State{util.Point{0, 0}, 0}
	visited := util.NewSet[util.Point]()
	visited.Add(start.pos)

	queue := make([]State, 0, 100)
	queue = append(queue, start)
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if s.pos.X == width-1 && s.pos.Y == height-1 {
			return s.steps, true
		}

		for _, n := range getNeighbors(m, s) {
			if visited.Has(n.pos) {
				continue
			}
			visited.Add(n.pos)
			queue = append(queue, n)
		}
	}

	return 0, false
}

func getSlice(bytes [][]int, example bool) [][]int {
	if example {
		return bytes[:12]
	}
	return bytes[:1024]
}

func getNeighbors(m [][]string, state State) []State {
	states := make([]State, 0)
	deltas := []util.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, delta := range deltas {
		x, y := state.pos.X+delta.X, state.pos.Y+delta.Y
		if x >= 0 && y >= 0 && y < len(m) && x < len(m[y]) && m[y][x] != "#" {
			states = append(states, State{util.Point{x, y}, state.steps + 1})
		}
	}
	return states
}

type State struct {
	pos   util.Point
	steps int
}

func getSize(example bool) (width, height int) {
	if example {
		return 7, 7
	}
	return 71, 71
}

func printMap(m [][]string) {
	for _, row := range m {
		for _, cell := range row {
			if cell == "#" {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func coordsToMap(coords [][]int, width, height int) [][]string {
	m := util.Create2DArray[string](width, height)
	for _, p := range coords {
		m[p[1]][p[0]] = "#"
	}
	return m
}
