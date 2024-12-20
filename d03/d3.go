package d03

import (
	"awesomeProject/util"
	"bufio"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 3,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "161", Real: "185797128"},
		},
		Part2: util.Solution{
			Solver:       Part2,
			Expected:     util.Expected{Example: "48", Real: "89798695"},
			ExampleInput: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
		},
	}
}

func Part1(input string, example bool) string {
	sum := 0
	for _, pair := range parse(input, false) {
		sum += pair[0] * pair[1]
	}

	//sum := 0
	//tokens := tokenize(input)
	//fmt.Println(tokens)
	//for _, token := range tokens {
	//	parts := strings.Split(token, ",")
	//	left, err := strconv.Atoi(parts[0])
	//	if err != nil {
	//		return "error"
	//	}
	//	right, err := strconv.Atoi(parts[1])
	//	if err != nil {
	//		return "error"
	//	}
	//	sum += left * right
	//}

	return strconv.Itoa(sum)
}

func tokenize(str string) []string {
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)
	scanner.Split(split)

	tokens := make([]string, 0)
	for scanner.Scan() {
		token := scanner.Bytes()
		if string(token) == "" {
			continue
		}

		tokens = append(tokens, string(token))
		//fmt.Printf("Found token %q\n", string(token), ",")
	}

	return tokens
}

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//fmt.Println("SPLIT", string(data[:]), string(data[:4]))
	if len(data) >= 4 && string(data[:4]) == "mul(" {
		return consumeMul(data)
	} else {
		return consumeUntil(data)
	}
}

func consumeMul(data []byte) (int, []byte, error) {
	var accum []byte
	stage := 0
	digits := 0
	for i, b := range data[4:] {
		if stage == 0 && util.IsDigit(rune(b)) {
			accum = append(accum, b)
			digits++
		} else if stage == 0 && b == ',' {
			accum = append(accum, b)
			stage++
			digits = 0
		} else if stage == 1 && util.IsDigit(rune(b)) {
			accum = append(accum, b)
			digits++
		} else if stage == 1 && b == ')' {
			return 3 + i, accum, nil
		} else {
			return 3 + i, []byte(""), nil
		}

		if digits > 3 {
			return 3 + i, []byte(""), nil
		}
	}
	return 3 + len(accum), []byte(""), nil

}

func consumeUntil(data []byte) (int, []byte, error) {
	var accum []byte
	for i, b := range data {
		if b != 'm' {
			accum = append(accum, b)
		} else {
			if i == 0 {
				return 1, []byte(""), nil
			}
			return i, []byte(""), nil
		}
	}
	if len(accum) == 0 {
		return 1, []byte(""), nil
	}
	return len(accum), []byte(""), nil
}

func consumeInt(data []byte) (int, []byte, error) {
	var accum []byte
	for i, b := range data {
		if util.IsDigit(rune(b)) {
			accum = append(accum, b)
		} else {
			return i, accum, nil
		}
	}
	return len(accum), accum, nil
}

func parse(str string, doDo bool) [][]int {
	pairs := make([][]int, 0)

	do := true
	for i := 0; i < len(str); {

		if i+7 < len(str) && str[i:i+7] == "don't()" {
			i += 7
			do = false
		} else if i+4 < len(str) && str[i:i+4] == "do()" {
			i += 4
			do = true
		}

		next, nums := parseMult(str, i)
		i = next
		if len(nums) == 0 {
			continue
		} else if !doDo || do {
			pairs = append(pairs, nums)
		}
	}

	return pairs
}

func parseMult(str string, i int) (int, []int) {
	nums := make([]int, 0)
	if str[i] != 'm' {
		return i + 1, nums
	} else if str[i+1] != 'u' {
		return i + 2, nums
	} else if str[i+2] != 'l' {
		return i + 3, nums
	} else if str[i+3] != '(' {
		return i + 4, nums
	}

	i, left, ok := util.ParseNumber(str, i+4)
	if !ok || left >= 1000 {
		return i, nums
	}

	if str[i] != ',' {
		return i + 1, nums
	}

	i, right, ok := util.ParseNumber(str, i+1)
	if !ok || right >= 1000 {
		return i, nums
	}

	if str[i] != ')' {
		return i + 1, nums
	}

	nums = append(nums, left, right)
	return i + 1, nums
}

func Part2(input string, example bool) string {
	sum := 0
	for _, pair := range parse(input, true) {
		sum += pair[0] * pair[1]
	}

	return strconv.Itoa(sum)
}
