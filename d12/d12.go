package d12

import (
	"awesomeProject/util"
	"strconv"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 12,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "1930", Real: "1477924"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "1206", Real: "841934"},
		},
	}
}

func Part1(input string, example bool) string {
	data := util.ParseCharMatrix(input)

	visited := util.NewSet[util.Point]()
	sum := 0
	for y, row := range data {
		for x := range row {
			p := util.Point{x, y}
			if visited.Has(p) {
				continue
			}
			field := util.FloodFill(data, p, getNeighbors)
			visited.Combine(field)
			perimeter := fieldPerimeter(field)
			area := field.Size()
			sum += area * perimeter

		}
	}

	return strconv.Itoa(sum)
}

func Part2(input string, example bool) string {
	data := util.ParseCharMatrix(input)

	visited := util.NewSet[util.Point]()
	sum := 0
	for y, row := range data {
		for x := range row {
			p := util.Point{x, y}
			if visited.Has(p) {
				continue
			}
			field := util.FloodFill(data, p, getNeighbors)
			visited.Combine(field)
			sides := fieldSides(field)
			area := field.Size()
			sum += area * sides

		}
	}

	return strconv.Itoa(sum)
}

var deltas = []util.Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func fieldSides(field *util.Set[util.Point]) (sides int) {
	fieldList := field.ToList()
	min, max := util.Point{fieldList[0].X, fieldList[0].Y}, util.Point{fieldList[0].X, fieldList[0].Y}
	for p := range field.List {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			p := util.Point{x, y}
			p2 := util.Point{x, y - 1}
			if field.Has(p) && !field.Has(p2) {
				sides++
				for field.Has(p) && !field.Has(p2) {
					x++
					p = util.Point{x, y}
					p2 = util.Point{x, y - 1}
				}
			}
		}
		for x := min.X; x <= max.X; x++ {
			p := util.Point{x, y}
			p2 := util.Point{x, y + 1}
			if field.Has(p) && !field.Has(p2) {
				sides++
				for field.Has(p) && !field.Has(p2) {
					x++
					p = util.Point{x, y}
					p2 = util.Point{x, y + 1}
				}
			}
		}
	}

	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			p := util.Point{x, y}
			p2 := util.Point{x - 1, y}
			if field.Has(p) && !field.Has(p2) {
				sides++
				for field.Has(p) && !field.Has(p2) {
					y++
					p = util.Point{x, y}
					p2 = util.Point{x - 1, y}
				}
			}
		}
		for y := min.Y; y <= max.Y; y++ {
			p := util.Point{x, y}
			p2 := util.Point{x + 1, y}
			if field.Has(p) && !field.Has(p2) {
				sides++
				for field.Has(p) && !field.Has(p2) {
					y++
					p = util.Point{x, y}
					p2 = util.Point{x + 1, y}
				}
			}
		}
	}
	return
}

func fieldPerimeter(field *util.Set[util.Point]) (perimeter int) {
	for p := range field.List {
		for _, d := range deltas {
			x, y := p.X+d.X, p.Y+d.Y
			if !field.Has(util.Point{x, y}) {
				perimeter++
			}
		}
	}
	return
}

func getNeighbors(farm [][]string, p util.Point) (neighbors []util.Point) {
	crop := farm[p.Y][p.X]
	for _, delta := range deltas {
		x, y := p.X+delta.X, p.Y+delta.Y
		if util.InArr(farm, x, y) && farm[y][x] == crop {
			neighbors = append(neighbors, util.Point{x, y})
		}
	}
	return
}
