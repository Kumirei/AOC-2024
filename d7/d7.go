package d7

import (
	"awesomeProject/util"
	"math"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 7,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "3749", Real: "975671981569"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "11387", Real: "223472064194845"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)

	operators := []Operator{Add, Multiply}
	sum := 0
	for _, test := range data {
		ok := validateResult(test.Result, test.Values, operators)
		if ok {
			sum += test.Result
		}
	}

	return strconv.Itoa(sum)
}

func Part2(input string, example bool) string {
	data := parse(input)

	operators := []Operator{Add, Multiply, Concat}
	sum := 0
	for _, test := range data {
		ok := validateResult(test.Result, test.Values, operators)
		if ok {
			sum += test.Result
		}
	}

	return strconv.Itoa(sum)
}

func validateResult(result int, values []int, operators []Operator) bool {
	//fmt.Println("VALIDATE", result, values)
	if len(values) == 0 {
		return false
	}
	lastIndex := len(values) - 1
	next := values[lastIndex]
	if len(values) == 1 && result == next {
		return true
	}
	for _, op := range operators {
		nextResult, ok := inverseOperator(result, next, op)
		if ok && validateResult(nextResult, values[:lastIndex], operators) {
			return true
		}
	}
	return false
}

func inverseOperator(a, b int, operator Operator) (int, bool) {
	switch operator {
	case Add:
		return a - b, true
	case Multiply:
		if a%b == 0 {
			return a / b, true
		}
	case Concat:
		digits := util.CountDigits(b)
		if util.LastNDigits(a, digits) == b {
			divisor := int(math.Pow(10, float64(digits)))
			return a / divisor, true
		}
	}
	return 0, false
}

type Operator int

const (
	Add Operator = iota
	Multiply
	Concat
)

func parse(input string) (data Data) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result, _ := strconv.Atoi(parts[0])
		test := Test{}
		test.Result = result
		for _, num := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(num)
			test.Values = append(test.Values, n)
		}
		data = append(data, test)
	}
	return
}

type Data []Test

type Test struct {
	Result int
	Values []int
}
