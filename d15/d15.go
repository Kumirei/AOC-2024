package d15

import (
	"awesomeProject/util"
	"fmt"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 15,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "10092", Real: "1451928"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "9021", Real: "1462788"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	for _, m := range data.Moves {
		move(data, data.Robot, m)
	}

	return strconv.Itoa(score(data))
}

func Part2(input string, example bool) string {
	data := parse(input)
	data = scale(data)
	for _, m := range data.Moves {
		move(data, data.Robot, m)
	}
	return strconv.Itoa(score(data))
}

func score(data *Data) (sum int) {
	for k, v := range data.Map {
		if v != Box && v != LeftBox {
			continue
		}
		sum += k.X + k.Y*100
	}
	return
}

func draw(data *Data) {
	for y := range data.Height {
		for x := range data.Width {
			fmt.Print(data.Map[util.Point{X: x, Y: y}])
		}
		fmt.Println()
	}
}

func move(data *Data, pos *util.Point, m Move) {
	delta := util.Point{}
	switch m {
	case Up:
		delta.Y = -1
	case Right:
		delta.X = 1
	case Left:
		delta.X = -1
	case Down:
		delta.Y = 1
	}

	ok := canMove(data, pos, &delta, false)
	if ok {
		moveAll(data, pos, &delta, false)
	}
}

func moveAll(data *Data, pos *util.Point, delta *util.Point, skipBoxCheck bool) {
	next := util.Point{X: pos.X + delta.X, Y: pos.Y + delta.Y}
	nextType := data.Map[next]

	if nextType != Floor {
		moveAll(data, &next, delta, false)
	}

	t := data.Map[*pos]
	if !skipBoxCheck && t == LeftBox && (delta.Y == 1 || delta.Y == -1) {
		moveAll(data, &util.Point{X: pos.X + 1, Y: pos.Y}, delta, true)
	} else if !skipBoxCheck && t == RightBox && (delta.Y == 1 || delta.Y == -1) {
		moveAll(data, &util.Point{X: pos.X - 1, Y: pos.Y}, delta, true)
	}
	if t == Robot {
		data.Robot = &next
	}

	data.Map[next] = data.Map[*pos]
	data.Map[*pos] = Floor
}

func canMove(data *Data, pos *util.Point, delta *util.Point, skipBoxCheck bool) bool {
	next := util.Point{X: pos.X + delta.X, Y: pos.Y + delta.Y}
	nextType := data.Map[next]
	if nextType == Wall {
		return false
	}

	t := data.Map[*pos]
	ok := true
	if !skipBoxCheck && t == LeftBox && (delta.Y == 1 || delta.Y == -1) {
		ok = canMove(data, &util.Point{X: pos.X + 1, Y: pos.Y}, delta, true)
	} else if !skipBoxCheck && t == RightBox && (delta.Y == 1 || delta.Y == -1) {
		ok = canMove(data, &util.Point{X: pos.X - 1, Y: pos.Y}, delta, true)
	}
	if !ok {
		return false
	}

	if nextType == Floor || canMove(data, &next, delta, false) {
		return true
	}
	return false
}

func scale(data *Data) *Data {
	newData := Data{
		Width:  data.Width * 2,
		Height: data.Height,
		Robot:  &util.Point{X: data.Robot.X * 2, Y: data.Robot.Y},
		Map:    make(Map),
		Moves:  data.Moves,
	}

	for k, v := range data.Map {
		pos1 := util.Point{X: k.X * 2, Y: k.Y}
		pos2 := util.Point{X: k.X*2 + 1, Y: k.Y}
		newData.Map[pos1] = v
		if v == Robot {
			newData.Map[pos2] = Floor
		} else if v == Box {
			newData.Map[pos1] = LeftBox
			newData.Map[pos2] = RightBox
		} else {
			newData.Map[pos2] = v
		}
	}

	return &newData
}

func parse(input string) *Data {
	parts := strings.Split(input, "\n\n")

	data := Data{Map: make(Map), Moves: make([]Move, 0)}
	arr := util.ParseCharMatrix(parts[0])
	data.Height = len(arr)
	data.Width = len(arr[0])
	for y, line := range arr {
		for x, cell := range line {
			tile := parseTile(cell)
			pos := util.Point{X: x, Y: y}
			data.Map[pos] = tile
			if tile == Robot {
				data.Robot = &pos
			}
		}
	}

	for _, line := range strings.Split(parts[1], "\n") {
		for _, cell := range strings.Split(line, "") {
			data.Moves = append(data.Moves, parseMove(cell))
		}
	}

	return &data
}

func parseMove(char string) Move {
	switch char {
	case "<":
		return Left
	case "^":
		return Up
	case ">":
		return Right
	case "v":
		return Down
	}
	panic("Invalid move")
}

func parseTile(char string) Tile {
	switch char {
	case "#":
		return Wall
	case ".":
		return Floor
	case "O":
		return Box
	case "@":
		return Robot
	}
	return Unknown
}

type Data struct {
	Map    Map
	Moves  []Move
	Robot  *util.Point
	Width  int
	Height int
}

type Map map[util.Point]Tile

type Tile int

const (
	Wall Tile = iota
	Floor
	Box
	Robot
	Unknown
	LeftBox
	RightBox
)

func (tile Tile) String() string {
	switch tile {
	case Wall:
		return "#"
	case Floor:
		return "."
	case Box:
		return "O"
	case LeftBox:
		return "["
	case RightBox:
		return "]"
	case Robot:
		return "@"
	default:
		return "?"
	}
}

type Move int

const (
	Left Move = iota
	Up
	Right
	Down
)

func (move Move) String() string {
	switch move {
	case Left:
		return "<"
	case Right:
		return ">"
	case Up:
		return "^"
	case Down:
		return "v"
	default:
		return "?"
	}
}
