package d11

import (
	"awesomeProject/util"
	"strconv"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 11,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "55312", Real: "211306"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "65601038650482", Real: "250783680217283"},
		},
	}
}

func Part1(input string, example bool) string {
	stones := parse(input)
	stones = blinks(stones, 25)
	count := countStones(stones)
	return strconv.Itoa(count)
}

func Part2(input string, example bool) string {
	stones := parse(input)
	stones = blinks(stones, 75)
	count := countStones(stones)
	return strconv.Itoa(count)
}

func countStones(stones Stones) (count int) {
	for _, c := range stones {
		count += c
	}
	return
}

type Stones map[int]int // Map from val to count

func blinks(stones Stones, count int) Stones {
	for range count {
		stones = blink(stones)
	}
	return stones
}

func blink(stones Stones) Stones {
	newStones := make(Stones)
	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
			continue
		}
		digits := util.CountDigits(stone)
		if digits%2 == 0 {
			left := util.FirstNDigits(stone, digits/2)
			right := util.LastNDigits(stone, digits/2)
			newStones[left] += count
			newStones[right] += count
		} else {
			newStones[stone*2024] += count
		}
	}
	return newStones
}

func parse(input string) Stones {
	arr := util.ParseIntMatrix(input, " ")[0]
	stones := make(Stones)
	for _, stone := range arr {
		count := stones[stone]
		stones[stone] = count + 1
	}
	return stones
}
