package d1

import (
	"awesomeProject/util"
	"sort"
	"strconv"
)

type Data struct {
	left  []int
	right []int
}

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 1,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "11", Real: "1646452"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "31", Real: "23609874"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)

	sort.Slice(data.left, func(i, j int) bool { return data.left[i] < data.left[j] })
	sort.Slice(data.right, func(i, j int) bool { return data.right[i] < data.right[j] })

	return strconv.Itoa(totalDistance(data))
}

func Part2(input string, example bool) string {
	data := parse(input)
	return strconv.Itoa(similarity(data))
}

func similarity(data Data) int {
	right := frequency(data.right)

	score := 0
	for _, k := range data.left {
		r, ok := right[k]
		if !ok {
			r = 0
		}
		score += k * r
	}

	return score
}

func frequency(nums []int) map[int]int {
	m := make(map[int]int)
	for _, num := range nums {
		freq, ok := m[num]
		if !ok {
			freq = 0
		}
		m[num] = freq + 1
	}
	return m
}

func totalDistance(data Data) int {
	sum := 0
	for i, l := range data.left {
		r := data.right[i]
		diff := l - r
		if diff < 0 {
			sum += -diff
		} else {
			sum += diff
		}
	}
	return sum
}

func parse(input string) Data {
	arr := util.Parse2DIntArray(input, "   ")
	left := make([]int, 0)
	right := make([]int, 0)
	for _, row := range arr {
		left = append(left, row[0])
		right = append(right, row[1])
	}
	return Data{left, right}
}
