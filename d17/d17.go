package d17

import (
	"awesomeProject/util"
	"fmt"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 17,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "4,6,3,5,6,3,5,2,1,0", Real: "6,1,6,4,2,4,7,3,5"},
		},
		Part2: util.Solution{
			Solver:       Part2,
			Expected:     util.Expected{Example: "117440", Real: "202975183645226"},
			ExampleInput: "Register A: 2024\nRegister B: 0\nRegister C: 0\n\nProgram: 0,3,5,4,3,0",
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	result := strings.Join(util.MapList(data.run(), func(n int) string { return fmt.Sprint(n) }), ",")
	return result
}

func Part2(input string, example bool) string {
	data := parse(input)
	result, _ := solve(data.program, len(data.program)-1, 0)
	return strconv.Itoa(result)
}

func solve(program []int, position, A int) (int, bool) {
	if position < 0 {
		return A, true
	}
	var output int
	for o := range 8 {
		a := A<<3 | o
		b := 0
		c := 0
		i := 0
		for i < len(program) {
			instruction := program[i]
			operand := program[i+1]

			combo := operand
			switch operand {
			case 4:
				combo = a
			case 5:
				combo = b
			case 6:
				combo = c
			}

			if instruction == 0 {
				a >>= combo
			} else if instruction == 1 {
				b ^= operand
			} else if instruction == 2 {
				b = combo & 7
			} else if instruction == 3 {
				if a != 0 {
					i = operand - 2
				}
			} else if instruction == 4 {
				b ^= c
			} else if instruction == 5 {
				output = combo & 7
				break
			} else if instruction == 6 {
				b = a >> combo
			} else if instruction == 7 {
				c = a >> combo
			}
			i += 2
		}
		if output == program[position] {
			newA, ok := solve(program, position-1, A<<3|o)
			if ok {
				return newA, ok
			}
		}
	}
	return 0, false
}

func (data *Data) run() []int {
	i := 0
	output := make([]int, 0)
	for i < len(data.program) {
		instruction := data.program[i]
		operand := data.program[i+1]
		switch instruction {
		case adv: // A / combo
			data.A = data.A / util.PowInt(2, data.getCombo(operand))
		case bxl: // B XOR literal
			data.B = data.B ^ operand
		case bst: // combo % 8
			data.B = data.getCombo(operand) % 8
		case jnz: // Nothing if A=0, else jump to literal (don't advance)
			if data.A != 0 {
				i = operand
				continue
			}
		case bxc: // B XOR C
			data.B = data.B ^ data.C
		case out: // Output combo % 8
			output = append(output, data.getCombo(operand)%8)
		case bdv: // A / combo
			data.B = data.A / util.PowInt(2, data.getCombo(operand))
		case cdv: // A / combo
			data.C = data.A / util.PowInt(2, data.getCombo(operand))

		}
		i += 2
	}

	return output
}

func (data *Data) getCombo(operand int) int {
	switch operand {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return operand
	case 4:
		return data.A
	case 5:
		return data.B
	case 6:
		return data.C
	case 7:
		fallthrough
	default:
		panic("INVALID COMBO OPERAND")
	}
}

func parse(input string) Data {
	parts := strings.Split(input, "\n\n")
	var A, B, C int
	var s string
	fmt.Sscanf(parts[0], "Register %s %d\nRegister %s %d\nRegister %s %d", &s, &A, &s, &B, &s, &C)
	//fmt.Println(A, B, C, s)
	programStr := strings.Split(parts[1], " ")[1]
	program := util.Parse2DIntArray(programStr, ",")[0]
	return Data{A, B, C, program}
}

type Data struct {
	A       int
	B       int
	C       int
	program []int
}

const (
	adv int = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)
