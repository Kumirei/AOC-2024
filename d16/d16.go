package d16

import (
	"awesomeProject/util"
	"fmt"
	"strconv"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 16,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "11048", Real: "72400"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "64", Real: "435"},
		},
	}
}

const INT_MAX = int(^uint(0) >> 1)

func Part1(input string, example bool) string {
	data := util.ParseCharMatrix(input)
	start, end := findEnds(data)

	costs, _ := findBestPath(data, start, end)
	cost := min(costs[end.Y][end.X][0], costs[end.Y][end.X][1], costs[end.Y][end.X][2], costs[end.Y][end.X][3])

	return strconv.Itoa(cost)
}

func Part2(input string, example bool) string {
	data := util.ParseCharMatrix(input)
	start, end := findEnds(data)

	costs, bp := findBestPath(data, start, end)
	minCost := INT_MAX
	var minCostBP *util.Set[util.Point]
	for dir := range 4 {
		cost := costs[end.Y][end.X][dir]
		if cost < minCost {
			minCost = cost
			minCostBP = bp[Position{util.Point{end.X, end.Y}, dir}]
		}
	}
	//drawBestPaths(data, minCostBP)
	//printCosts(costs)
	return strconv.Itoa(minCostBP.Size())
}

func getMinCost(costs [4]int) (cost int) {
	return min(costs[0], costs[1], costs[2], costs[3])
}

func drawBestPaths(data [][]string, bp *util.Set[util.Point]) {
	for y, line := range data {
		for x, cell := range line {
			if bp.Has(util.Point{x, y}) {
				fmt.Print("O")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func findBestPath(data [][]string, start, end util.Point) ([][][4]int, map[Position]*util.Set[util.Point]) {
	costs := make([][][4]int, len(data))
	for y := range len(data) {
		costs[y] = make([][4]int, len(data[y]))
		for x := range len(data[y]) {
			costs[y][x] = [4]int{INT_MAX, INT_MAX, INT_MAX, INT_MAX}
		}
	}
	costs[start.Y][start.X][RIGHT] = 0

	startPosition := Position{start, RIGHT}

	bestPaths := make(map[Position]*util.Set[util.Point])
	bestPaths[startPosition] = util.NewSet[util.Point]()
	bestPaths[startPosition].Add(start)

	queue := []Position{startPosition}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		// Get cost of current node
		currentCost := costs[p.Pos.Y][p.Pos.X][p.Dir]
		currentBP := bestPaths[p]

		// Get neighbors
		neighbors := getNeighbors(data, p)
		for _, n := range neighbors {
			neighborCost := currentCost + 1 + (util.AbsInt(p.Dir-n.Dir)%2)*1000
			if neighborCost < costs[n.Pos.Y][n.Pos.X][n.Dir] {
				costs[n.Pos.Y][n.Pos.X][n.Dir] = neighborCost
				queue = append(queue, n)
				// Replace best path
				neighborBP := util.NewSet[util.Point]()
				neighborBP.Combine(currentBP)
				neighborBP.Add(n.Pos)
				bestPaths[n] = neighborBP

			} else if neighborCost == costs[n.Pos.Y][n.Pos.X][n.Dir] {
				// Add to best path
				neighborBP := bestPaths[n]
				if neighborBP == nil {
					neighborBP = util.NewSet[util.Point]()
				}
				neighborBP.Combine(currentBP)
			}
		}
	}

	return costs, bestPaths
}

func printCosts(costs [][][4]int) {
	for y := range costs {
		for x := range costs[y] {
			minCost := min(costs[y][x][0], costs[y][x][1], costs[y][x][2], costs[y][x][3])
			if minCost == INT_MAX {
				fmt.Print("###### ")
			} else {
				fmt.Printf("%6d ", minCost)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getNeighbors(data [][]string, pos Position) []Position {
	neighbors := make([]Position, 0)

	up := Position{util.Point{pos.Pos.X, pos.Pos.Y - 1}, UP}
	right := Position{util.Point{pos.Pos.X + 1, pos.Pos.Y}, RIGHT}
	down := Position{util.Point{pos.Pos.X, pos.Pos.Y + 1}, DOWN}
	left := Position{util.Point{pos.Pos.X - 1, pos.Pos.Y}, LEFT}
	if up.Pos.Y >= 0 && up.Pos.Y < len(data) && up.Pos.X >= 0 && up.Pos.X < len(data[up.Pos.Y]) && data[up.Pos.Y][up.Pos.X] != "#" {
		neighbors = append(neighbors, up)
	}
	if right.Pos.Y >= 0 && right.Pos.Y < len(data) && right.Pos.X >= 0 && right.Pos.X < len(data[right.Pos.Y]) && data[right.Pos.Y][right.Pos.X] != "#" {
		neighbors = append(neighbors, right)
	}
	if down.Pos.Y >= 0 && down.Pos.Y < len(data) && down.Pos.X >= 0 && down.Pos.X < len(data[down.Pos.Y]) && data[down.Pos.Y][down.Pos.X] != "#" {
		neighbors = append(neighbors, down)
	}
	if left.Pos.Y >= 0 && left.Pos.Y < len(data) && left.Pos.X >= 0 && left.Pos.X < len(data[left.Pos.Y]) && data[left.Pos.Y][left.Pos.X] != "#" {
		neighbors = append(neighbors, left)
	}

	return neighbors
}

type Position struct {
	Pos util.Point
	Dir int
}

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

func findEnds(data [][]string) (start, end util.Point) {
	for y, line := range data {
		for x, cell := range line {
			if cell == "S" {
				start = util.Point{x, y}
			} else if cell == "E" {
				end = util.Point{x, y}
			}
		}
	}
	return
}
