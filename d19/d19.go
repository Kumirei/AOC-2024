package d19

import (
	"awesomeProject/util"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 19,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "6", Real: "283"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "16", Real: "615388132411142"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	possible := 0
	for _, design := range data.designs {
		solutions := countSolutions(data.towels, design)
		if solutions > 0 {
			possible++
		}
	}
	return strconv.Itoa(possible)
}

func Part2(input string, example bool) string {
	data := parse(input)
	count := 0
	for _, design := range data.designs {
		solutions := countSolutions(data.towels, design)
		count += solutions
	}
	return strconv.Itoa(count)
}

func countSolutions(towels []string, design string) int {
	dp := make([]int, len(design)+1)
	dp[0] = 1
	towelSet := util.NewSet[string]()
	towelSet.AddMulti(towels...)
	for start := range len(design) {
		next := make([]int, len(design)+1)
		for end := start + 1; end <= len(design); end++ {
			subStr := design[start:end]
			count := dp[end]
			if towelSet.Has(subStr) {
				count += dp[start]
			}
			next[end] = count
		}
		dp = next
	}
	return dp[len(dp)-1]
}

func parse(input string) *Data {
	parts := strings.Split(input, "\n\n")
	data := Data{
		towels:  strings.Split(parts[0], ", "),
		designs: strings.Split(parts[1], "\n"),
	}

	return &data
}

type Data struct {
	towels  []string
	designs []string
}
