package d20

import (
	"awesomeProject/util"
	"strconv"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 20,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "44", Real: "1360"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "285", Real: "1005476"},
		},
	}
}

func Part1(input string, example bool) string {
	limit := 100
	if example {
		limit = 2
	}
	data := util.Parse2DCharArray(input)

	start := findStart(data)
	costs := getCosts(data, start)
	cheats := findCheats(costs, start, 2)
	count := countCheats(cheats, limit)

	return strconv.Itoa(count)
}

func Part2(input string, example bool) string {
	limit := 100
	if example {
		limit = 50
	}
	data := util.Parse2DCharArray(input)

	start := findStart(data)
	costs := getCosts(data, start)
	cheats := findCheats(costs, start, 20)
	count := countCheats(cheats, limit)

	return strconv.Itoa(count)
}

func countCheats(cheatMap map[int]int, limit int) int {
	count := 0
	for k, v := range cheatMap {
		if k >= limit {
			count += v
		}
	}
	return count
}

func findCheats(costs [][]int, start util.Point, allowedCheats int) map[int]int {
	cheatMap := make(map[int]int) // steps to count
	for _, p := range util.PointsOfArr(costs) {
		if costs[p.Y][p.X] == 0 && p != start {
			continue
		}
		cheats := cheat(costs, p, allowedCheats)
		for _, c := range cheats {
			cheatMap[c]++
		}
	}
	return cheatMap
}

func getCosts(data [][]string, start util.Point) [][]int {
	node := start
	costs := util.Create2DArray[int](len(data[0]), len(data))
	visited := util.NewSet[util.Point]()
	done := false
	for !done {
		cost := costs[node.Y][node.X]
		for _, n := range util.GetCardinalNeighborsInside(node, 0, len(data[0]), 0, len(data)) {
			if visited.Has(n) || data[n.Y][n.X] == "#" {
				continue
			}
			costs[n.Y][n.X] = cost + 1
			visited.Add(n)
			node = n
			if data[n.Y][n.X] == "E" {
				done = true
			}
			break
		}
	}
	return costs
}

func findStart(data [][]string) util.Point {
	var start util.Point
	for y, line := range data {
		for x, cell := range line {
			if cell == "S" {
				start = util.Point{x, y}
			}
		}
	}
	return start
}

func cheat(costs [][]int, p util.Point, allowedCheats int) []int {
	cheats := make([]int, 0)
	x, y := p.X, p.Y
	cost := costs[y][x]

	cheatPoints := getCheatPoints(costs, p, allowedCheats)

	for _, pos := range cheatPoints {
		posCost := costs[pos.Y][pos.X]
		cheatSteps := util.AbsInt(p.X-pos.X) + util.AbsInt(p.Y-pos.Y)
		if posCost > cost && posCost-cost-cheatSteps > 0 {
			cheats = append(cheats, posCost-cost-cheatSteps)
		}
	}
	return cheats
}

func getCheatPoints(costs [][]int, point util.Point, allowedCheats int) []util.Point {
	points := make([]util.Point, 0)
	for yi := range allowedCheats*2 + 1 {
		dy := yi - allowedCheats
		for xi := range allowedCheats*2 + 1 {
			dx := xi - allowedCheats
			if util.AbsInt(dy)+util.AbsInt(dx) > allowedCheats || (dx == 0 && dy == 0) {
				continue
			}
			p := util.Point{point.X + dx, point.Y + dy}
			if p.Y < 0 || p.X < 0 || p.Y >= len(costs) || p.X >= len(costs[p.Y]) {
				continue
			}
			points = append(points, p)
		}
	}
	return points
}
