package d22

import (
	"awesomeProject/util"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 22,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "37327623", Real: "20506453102"},
		},
		Part2: util.Solution{
			Solver:       Part2,
			Expected:     util.Expected{Example: "23", Real: "2423"},
			ExampleInput: "1\n2\n3\n2024\n",
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)

	sum := 0
	for _, start := range data {
		secret := Secret(start)
		for range 2000 {
			secret.next()
		}
		sum += int(secret)
	}

	return strconv.Itoa(sum)
}

func Part2(input string, example bool) string {
	data := parse(input)

	rounds := 2000

	bestSequences := make(map[string]int)

	for _, start := range data {
		secret := Secret(start)
		prices := make([]int, rounds)
		changes := make([]int, rounds-1)
		prices[0] = start % 10

		for i := range rounds - 1 {
			secret.next()
			price := int(secret) % 10
			prices[i+1] = price
			changes[i] = price - prices[i]
		}

		left := 0
		width := 4
		sequences := make(map[string]int)
		for left < rounds-width {
			right := left + width
			key := strings.Join(util.MapList(changes[left:right], util.ToString), ",")
			_, ok := sequences[key]
			if !ok {
				sequences[key] = prices[right]
				bestSequences[key] += prices[right]
			}
			left++
		}
	}

	bestSequence := 0
	for _, sum := range bestSequences {
		if sum > bestSequence {
			bestSequence = sum
		}
	}

	return strconv.Itoa(bestSequence)
}

type Secret int

func (s *Secret) next() *Secret {
	s.mix(int(*s) << 6).prune()
	s.mix(int(*s) >> 5).prune()
	s.mix(int(*s) << 11).prune()
	return s
}

func (s *Secret) mix(n int) *Secret {
	*s = Secret(int(*s) ^ n)
	return s
}

func (s *Secret) prune() *Secret {
	*s = Secret(int(*s) & ((1 << 24) - 1))
	return s
}

func parse(input string) []int {
	return util.ParseIntArray(input, "\n")
}
