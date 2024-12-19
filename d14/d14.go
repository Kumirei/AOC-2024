package d14

import (
	"awesomeProject/util"
	"fmt"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 14,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "12", Real: "232589280"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "none", Real: "7569"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	tick(data, 100, example)
	//draw(data, example)
	q1, q2, q3, q4 := countQuadrants(data, example)
	prod := q1 * q2 * q3 * q4

	return strconv.Itoa(prod)
}

func Part2(input string, example bool) string {
	if example {
		return "none"
	}
	data := parse(input)

	min := 999999999
	mini := 0
	for i := range 10000 {
		tick(data, 1, example)
		q1, q2, q3, q4 := countQuadrants(data, example)
		prod := q1 * q2 * q3 * q4
		if prod < min {
			min = prod
			mini = i + 1
		}
	}

	//data = parse(input)
	//tick(data, mini, example)
	//draw(data, example)

	return strconv.Itoa(mini)
}

func countQuadrants(robots []*Robot, example bool) (q1, q2, q3, q4 int) {
	width, height := getSize(example)
	X := width / 2
	Y := height / 2
	for _, robot := range robots {
		if robot.Pos.X < X && robot.Pos.Y < Y {
			q1++
		} else if robot.Pos.X > X && robot.Pos.Y < Y {
			q2++
		} else if robot.Pos.X < X && robot.Pos.Y > Y {
			q3++
		} else if robot.Pos.X > X && robot.Pos.Y > Y {
			q4++
		}
	}
	return
}

func tick(robots []*Robot, ticks int, example bool) {
	width, height := getSize(example)
	for _, robot := range robots {
		//fmt.Println(i, robot)
		robot.Pos.X = util.PosMod(robot.Pos.X+robot.Vel.X*ticks, width)
		robot.Pos.Y = util.PosMod(robot.Pos.Y+robot.Vel.Y*ticks, height)
		//fmt.Println(i, robot)
	}
}

func parse(input string) (robots []*Robot) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, &Robot{Pos: util.Point{X: px, Y: py}, Vel: util.Point{X: vx, Y: vy}})
	}

	return
}

func getSize(example bool) (width, height int) {
	if example {
		return 11, 7
	}
	return 101, 103
}

func draw(robots []*Robot, example bool) {
	width, height := getSize(example)
	arr := util.Create2DArray[int](width, height)

	for _, robot := range robots {
		arr[robot.Pos.Y%height][robot.Pos.X%width]++
	}

	//fmt.Println()
	for _, row := range arr {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
}

type Robot struct {
	Pos util.Point
	Vel util.Point
}

func (robot Robot) String() string {
	return fmt.Sprintf("Robot(Pos%v Vel%v)", robot.Pos, robot.Vel)
}
