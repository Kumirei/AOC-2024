package d13

import (
	"awesomeProject/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 13,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "480", Real: "33427"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "875318608908", Real: "91649162972270"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)

	tokens := 0
	for _, claw := range data {
		A, B, ok := solveEquation(claw)

		if ok {
			tokens += A*3 + B*1
		}
	}

	return strconv.Itoa(tokens)
}

func Part2(input string, example bool) string {
	data := parse(input)

	tokens := 0
	for _, claw := range data {
		claw.Prize.X += 10000000000000
		claw.Prize.Y += 10000000000000
		A, B, ok := solveEquation(claw)

		if ok {
			tokens += A*3 + B*1
		}

	}

	// 69953912952434 LOW
	// 90396486156559 LOW
	// 91649162972270 RIGHT!

	return strconv.Itoa(tokens)
}

func solveEquation(eq Claw) (A int, B int, ok bool) {
	frac := eq.B.Step.Y / eq.B.Step.X
	an := eq.Prize.Y - frac*eq.Prize.X
	ad := eq.A.Step.Y - frac*eq.A.Step.X
	a := an / ad

	bn := eq.Prize.X - eq.A.Step.X*a
	bd := eq.B.Step.X
	b := bn / bd

	if !isInteger(a) || !isInteger(b) {
		return 0, 0, false
	}
	return int(math.Round(a)), int(math.Round(b)), true
}

type Point struct {
	X float64
	Y float64
}

type Claw struct {
	A     Button
	B     Button
	Prize Point
}

type Button struct {
	Cost int
	Step Point
}

func parse(input string) (claws []Claw) {
	clawsStr := strings.Split(input, "\n\n")
	for _, clawStr := range clawsStr {
		lines := strings.Split(clawStr, "\n")
		var values []Point
		for i, line := range lines {
			var x, y float64
			var s string
			if i < 2 {
				fmt.Sscanf(line, "Button %s X+%f, Y+%f", &s, &x, &y)
			} else {
				fmt.Sscanf(line, "Prize: X=%f, Y=%f", &x, &y)
			}
			values = append(values, Point{X: x, Y: y})
		}
		claw := Claw{A: Button{Cost: 3, Step: values[0]}, B: Button{Cost: 1, Step: values[1]}, Prize: values[2]}
		claws = append(claws, claw)
	}
	return
}

func isInteger(n float64) bool {
	precision := 1e-3
	if _, frac := math.Modf(n); frac < precision || frac > 1-precision {
		return true
	}
	return false
}
