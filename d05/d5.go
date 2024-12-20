package d05

import (
	"awesomeProject/util"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 5,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "143", Real: "5166"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "123", Real: "4679"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)

	// Get correct reports
	correct := make([]Report, 0)
	for _, report := range data.Reports {
		if isCorrect(data.EdgeMap, report) {
			correct = append(correct, report)
		}
	}

	// Get values
	middles := make([]int, len(correct))
	for i, report := range correct {
		middles[i] = report[len(report)/2]
	}

	sum := util.Sum(middles)

	return strconv.Itoa(sum)
}

func Part2(input string, example bool) string {
	data := parse(input)

	// Get incorrect reports
	incorrect := make([]Report, 0)
	for _, report := range data.Reports {
		if !isCorrect(data.EdgeMap, report) {
			incorrect = append(incorrect, report)
		}
	}

	// Sort incorrect reports
	for _, report := range incorrect {
		reportSet := util.NewSet[int]()
		reportSet.AddMulti(report...)
		order := make(map[Node]int)
		for _, n := range report {
			edgeSet := data.EdgeMap[n]
			var intersect *util.Set[int]
			if edgeSet != nil {
				intersect = edgeSet.Intersect(reportSet)
				order[n] = len(report) - intersect.Size()
			} else {
				order[n] = len(report)
			}
		}
		sort.Slice(report, func(i, j int) bool { return order[report[i]] < order[report[j]] })
	}

	// Get values
	middles := make([]int, len(incorrect))
	for i, report := range incorrect {
		middles[i] = report[len(report)/2]
	}

	sum := util.Sum(middles)

	return strconv.Itoa(sum)
}

func countGreaterOnLeft(data Data, report Report, i int) int {
	greater := data.EdgeMap[report[i]]
	fmt.Println(report, report[i], greater)
	if greater == nil {
		return 0
	}
	count := 0
	for j := i - 1; j >= 0; j-- {
		if greater.Has(report[j]) {
			count++
		}
	}
	return count
}

func isCorrect(edges map[Node]*util.Set[int], report Report) bool {
	for i := len(report) - 1; i > 0; i-- {
		greater := edges[report[i]]
		if greater == nil {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if greater.Has(report[j]) {
				return false
			}
		}
	}
	return true
}

type Data struct {
	Edges   []Edge
	EdgeMap map[Node]*util.Set[int]
	Reports []Report
}

type Edge struct {
	From Node
	To   Node
}

type Report = []Node

type Node = int

func parse(str string) (data Data) {
	parts := strings.Split(str, "\n\n")
	if len(parts) != 2 {
		log.Fatal("Could not parse input data")
	}

	// Parse edges
	data.EdgeMap = make(map[Node]*util.Set[int])
	for _, line := range strings.Split(parts[0], "\n") {
		var left, right int
		fmt.Sscanf(line, "%d|%d", &left, &right)
		data.Edges = append(data.Edges, Edge{left, right})
		set, ok := data.EdgeMap[left]
		if !ok {
			set = util.NewSet[int]()
			data.EdgeMap[left] = set
		}
		set.Add(right)
	}

	// Parse reports
	for _, line := range strings.Split(parts[1], "\n") {
		nums := strings.Split(line, ",")
		report := make(Report, len(nums))
		for i, num := range nums {
			n, _ := strconv.Atoi(num)
			report[i] = n
		}
		data.Reports = append(data.Reports, report)
	}

	return
}
