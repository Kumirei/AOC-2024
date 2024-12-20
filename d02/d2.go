package d02

import (
	"awesomeProject/util"
	"strconv"
)

type Data = [][]int

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 2,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "2", Real: "606"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "4", Real: "644"},
		},
	}
}

func Part1(input string, example bool) string {
	data := util.Parse2DIntArray(input, " ")
	return strconv.Itoa(solve(data, false))
}

func Part2(input string, example bool) string {
	data := util.Parse2DIntArray(input, " ")
	return strconv.Itoa(solve(data, true))
}

func solve(data Data, tolerateOne bool) int {
	safe := 0
	for _, row := range data {
		if isSafe(row, tolerateOne) {
			safe++
		}
	}

	return safe

}

func isSafe(row []int, tolerateOne bool) bool {
	order := "none"
	for i := 1; i < len(row); i++ {
		diff := row[i] - row[i-1]

		if order == "none" {
			if diff > 0 {
				order = "asc"
			} else if diff < 0 {
				order = "desc"
			}
		}

		ok := isDiffSafe(order, diff)
		if !ok && !tolerateOne {
			return false
		} else if !ok && tolerateOne {
			omitPrev := util.CopySlice(row[:])
			omitPrev = util.Remove(omitPrev[:], i-1)
			if isSafe(omitPrev, false) {
				return true
			}
			omitCurr := util.CopySlice(row[:])
			omitCurr = util.Remove(omitCurr[:], i)
			if isSafe(omitCurr, false) {
				return true
			}
			if i == 2 {
				omitFirst := util.CopySlice(row[:])
				omitFirst = util.Remove(omitFirst[:], 0)
				return isSafe(omitFirst, false)
			}
			return false
		}
	}

	return true
}

func isDiffSafe(order string, diff int) bool {
	if diff == 0 || diff > 3 || diff < -3 {
		return false
	}
	if order == "asc" && diff < 0 {
		return false
	} else if order == "desc" && diff > 0 {
		return false
	}

	return true
}
