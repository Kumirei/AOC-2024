package template

import (
	"awesomeProject/util"
	"fmt"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 0,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "???", Real: "???"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "???", Real: "???"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	fmt.Println(data)
	return "TODO"
}

func Part2(input string, example bool) string {
	data := parse(input)
	fmt.Println(data)
	return "TODO"
}

func parse(input string) string {
	return "TODO"
}
