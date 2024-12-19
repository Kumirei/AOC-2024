package d6

import (
	"awesomeProject/util"
	"strconv"
	"strings"
)

const day = 6

func Solutions() util.Solutions {
	return util.Solutions{
		Day: day,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "41", Real: "5145"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "6", Real: "1523"},
		},
	}
}

func Part1(input string, example bool) string {
	m := parse(input)
	visited := walk(m)

	return strconv.Itoa(visited.Size())
}

func Part2(input string, example bool) string {
	m := parse(input)
	visited := walk(m)

	// Try setting every point to an obstacle and detect loops
	loops := 0
	for p := range visited.List {
		m.Obstacles.Add(p)
		if detectLoop(m) {
			loops++
		}
		m.Obstacles.Remove(p)
	}

	return strconv.Itoa(loops)
}

func detectLoop(m Map) bool {
	visited := util.NewSet[Guard]()
	guard := Guard{m.Guard.Pos, m.Guard.Dir}
	for inBounds(m, guard) {
		if visited.Has(guard) {
			return true
		}
		visited.Add(guard)
		step(m, &guard)
	}

	return false
}

func walk(m Map) *util.Set[Point] {
	visited := util.NewSet[Point]()
	guard := Guard{m.Guard.Pos, m.Guard.Dir}
	for inBounds(m, guard) {
		visited.Add(guard.Pos)
		step(m, &guard)
	}
	return visited
}

func inBounds(m Map, guard Guard) bool {
	x, y := guard.Pos.X, guard.Pos.Y
	return x >= 0 && y >= 0 && x < m.Width && y < m.Height
}

func step(m Map, guard *Guard) {
	var next Point
	var turn Direction
	switch guard.Dir {
	case up:
		next = Point{guard.Pos.X, guard.Pos.Y - 1}
		turn = right
	case right:
		next = Point{guard.Pos.X + 1, guard.Pos.Y}
		turn = down
	case down:
		next = Point{guard.Pos.X, guard.Pos.Y + 1}
		turn = left
	case left:
		next = Point{guard.Pos.X - 1, guard.Pos.Y}
		turn = up
	}
	if m.Obstacles.Has(next) {
		guard.Dir = turn
	} else {
		guard.Pos = next
	}
}

func parse(str string) (m Map) {
	m.Obstacles = util.NewSet[Point]()
	lines := strings.Split(str, "\n")
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				m.Obstacles.Add(Point{x, y})
			} else if char == '^' {
				m.Guard.Pos = Point{x, y}
				m.Guard.Dir = up
			}
		}
	}
	m.Height = len(lines)
	m.Width = len(lines[0])
	return
}

type Set[T comparable] util.Set[T]

type Point util.Point

type Direction int

const (
	up Direction = iota
	right
	down
	left
)

type Guard struct {
	Pos Point
	Dir Direction
}

type Map struct {
	Obstacles *util.Set[Point]
	Guard     Guard
	Width     int
	Height    int
}
