package d08

import (
	"awesomeProject/util"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 8,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "14", Real: "265"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "34", Real: "962"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	antiNodes := getAntiNodes(data.Antennas, data.Width, data.Height, 1)
	unique := uniqueAntiNodes(antiNodes, data.Width, data.Height)
	return strconv.Itoa(len(unique))
}

func Part2(input string, example bool) string {
	data := parse(input)
	antiNodes := getAntiNodes(data.Antennas, data.Width, data.Height, 2)
	unique := uniqueAntiNodes(antiNodes, data.Width, data.Height)
	return strconv.Itoa(len(unique))
}

func uniqueAntiNodes(antiMap AntiNodes, width, height int) (unique []util.Point) {
	set := util.NewSet[util.Point]()
	for _, antiNodes := range antiMap {
		for _, node := range antiNodes {
			if !util.InBounds(width, height, node.X, node.Y) {
				continue
			}
			set.Add(node)
		}
	}
	return set.ToList()
}

func getAntiNodes(antennas []Antenna, width, height int, part int) (antiMap AntiNodes) {
	antiMap = make(AntiNodes)
	for i, antenna := range antennas {
		antiNodes := antiMap[antenna.Freq]
		for j := i + 1; j < len(antennas); j++ {
			if antennas[j].Freq != antenna.Freq {
				continue
			}
			if part == 1 {
				antiNodes = append(antiNodes, createAntiNodes1(antenna.Position, antennas[j].Position)...)
			} else {
				antiNodes = append(antiNodes, createAntiNodes2(antenna.Position, antennas[j].Position, width, height)...)
			}
		}
		antiMap[antenna.Freq] = antiNodes
	}
	return
}

func createAntiNodes1(a, b util.Point) (antiNodes []util.Point) {
	dx, dy := a.X-b.X, a.Y-b.Y
	return []util.Point{{a.X + dx, a.Y + dy}, {b.X - dx, b.Y - dy}}
}

func createAntiNodes2(a, b util.Point, width, height int) (antiNodes []util.Point) {
	dx, dy := a.X-b.X, a.Y-b.Y

	x, y := a.X, a.Y
	// Go all the way to the edge
	for util.InBounds(width, height, x, y) {
		x, y = x+dx, y+dy
	}
	// Go all the way to the other edge
	x, y = x-dx, y-dy
	for util.InBounds(width, height, x, y) {
		antiNodes = append(antiNodes, util.Point{x, y})
		x, y = x-dx, y-dy
	}

	return
}

func parse(input string) (data Data) {
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, freq := range line {
			if freq == '.' {
				continue
			}
			data.Antennas = append(data.Antennas, Antenna{util.Point{x, y}, string(freq)})
		}
	}
	data.Height = len(lines)
	data.Width = len(lines[0])
	return data
}

type Data struct {
	Antennas []Antenna
	Width    int
	Height   int
}

type Antenna struct {
	Position util.Point
	Freq     string
}

type AntiNodes map[string][]util.Point
