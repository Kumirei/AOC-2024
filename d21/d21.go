package d21

import (
	"awesomeProject/util"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 21,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "126384", Real: "163086"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "154115708116294", Real: "198466286401228"},
		},
	}
}

func Part1(input string, example bool) string {
	data := util.ParseCharMatrix(input)

	totalComplexity := 0
	for _, row := range data {
		k := newKeypad()
		keys := k.arrowInputFor(row)
		memo := make(map[State]int)
		cost := costOf(State{strings.Join(keys, ""), 2}, memo)
		totalComplexity += complexityOfInput(row, cost)
	}

	return strconv.Itoa(totalComplexity)
}

func Part2(input string, example bool) string {
	data := util.ParseCharMatrix(input)

	totalComplexity := 0
	for _, row := range data {
		k := newKeypad()
		keys := k.arrowInputFor(row)
		memo := make(map[State]int)
		cost := costOf(State{strings.Join(keys, ""), 25}, memo)
		totalComplexity += complexityOfInput(row, cost)
	}

	return strconv.Itoa(totalComplexity)
}

func costOf(state State, memo map[State]int) int {
	cost, ok := memo[state]
	if ok {
		return cost
	}
	if state.depth == 0 {
		return len(state.key)
	}
	A := newArrowPad()
	cost = 0
	for _, char := range state.key {
		moves := A.goTo(string(char))
		moves = append(moves, "A")
		cost += costOf(State{strings.Join(moves, ""), state.depth - 1}, memo)
	}
	memo[state] = cost
	return cost
}

type State struct {
	key   string
	depth int
}

func complexityOfInput(input []string, cost int) int {
	code, _ := strconv.Atoi(strings.Join(util.MapList(input[:len(input)-1], util.ToString), ""))
	return cost * code
}

type ArrowPad util.Point

func newArrowPad() ArrowPad {
	return ArrowPad{2, 0}
}

func (A *ArrowPad) arrowInputFor(keys []string) []string {
	moves := make([]string, 0, len(keys))
	for _, key := range keys {
		newMoves := A.goTo(key)
		moves = append(moves, newMoves...)
		moves = append(moves, "A")
	}
	return moves
}

func (A *ArrowPad) goTo(key string) []string {
	moves := make([]string, 0)
	newCoords := A.coordsOf(key)
	dx, dy := newCoords.X-A.X, newCoords.Y-A.Y
	x, y := A.X+dx, A.Y+dy

	if x == 0 && y == 1 && A.Y == 0 {
		for range dy {
			moves = append(moves, "v")
			A.Y++
		}
		dy = 0
	}
	if A.X == 0 && A.Y == 1 && y == 0 {
		for range dx {
			moves = append(moves, ">")
			A.X++
		}
		dx = 0
	}

	if dx < 0 {
		for range -dx {
			moves = append(moves, "<")
			A.X--
		}
	}
	if dy > 0 {
		for range dy {
			moves = append(moves, "v")
			A.Y++
		}
	}
	if dy < 0 {
		for range -dy {
			moves = append(moves, "^")
			A.Y--
		}
	}
	if dx > 0 {
		for range dx {
			moves = append(moves, ">")
			A.X++
		}
	}

	return moves
}

func (A *ArrowPad) coordsOf(key string) util.Point {
	switch key {
	case "^":
		return util.Point{1, 0}
	case "A":
		return util.Point{2, 0}
	case "<":
		return util.Point{0, 1}
	case "v":
		return util.Point{1, 1}
	case ">":
		return util.Point{2, 1}
	default:
		panic("NO KEY")
	}
}

type Keypad util.Point

func newKeypad() Keypad {
	return Keypad{2, 3}
}

func (k *Keypad) arrowInputFor(keys []string) []string {
	moves := make([]string, 0, len(keys))
	for _, key := range keys {
		newMoves := k.goTo(key)
		moves = append(moves, newMoves...)
		moves = append(moves, "A")
	}
	return moves
}

func (k *Keypad) goTo(key string) []string {
	moves := make([]string, 0)
	newCoords := k.coordsOf(key)
	dx, dy := newCoords.X-k.X, newCoords.Y-k.Y
	x, y := k.X+dx, k.Y+dy

	if x == 0 && k.X != 0 && k.Y == 3 {
		for range -dy {
			moves = append(moves, "^")
			k.Y--
		}
		dy = 0
	}
	if k.X == 0 && x != 0 && y == 3 {
		for range dx {
			moves = append(moves, ">")
			k.X++
		}
		dx = 0
	}

	if dx < 0 {
		for range -dx {
			moves = append(moves, "<")
			k.X--
		}
	}
	if dy > 0 {
		for range dy {
			moves = append(moves, "v")
			k.Y++
		}
	}
	if dy < 0 {
		for range -dy {
			moves = append(moves, "^")
			k.Y--
		}
	}
	if dx > 0 {
		for range dx {
			moves = append(moves, ">")
			k.X++
		}
	}

	return moves
}

func (k *Keypad) coordsOf(key string) util.Point {
	switch key {
	case "9":
		return util.Point{2, 0}
	case "8":
		return util.Point{1, 0}
	case "7":
		return util.Point{0, 0}
	case "6":
		return util.Point{2, 1}
	case "5":
		return util.Point{1, 1}
	case "4":
		return util.Point{0, 1}
	case "3":
		return util.Point{2, 2}
	case "2":
		return util.Point{1, 2}
	case "1":
		return util.Point{0, 2}
	case "0":
		return util.Point{1, 3}
	case "A":
		return util.Point{2, 3}
	default:
		panic("NO KEY")
	}
}
