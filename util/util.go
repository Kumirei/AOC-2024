package util

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Inputs struct {
	Example string
	Real    string
}

func ReadInputs(day int) Inputs {
	example := ReadFile(day, true)
	input := ReadFile(day, false)
	return Inputs{example, input}
}

func ReadFile(day int, example bool) string {
	path := "d" + fmt.Sprintf("%02d", day) + "/"
	if example {
		path += "example"
	} else {
		path += "real"
	}
	path += ".txt"
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	//return string(data)
	//fmt.Println(string(data), strings.TrimSpace(string(data)))
	return strings.TrimSpace(string(data))
}

func ParseIntMatrix(input string, delimiter string) [][]int {
	parser := func(str string) (int, bool) {
		n, err := strconv.Atoi(str)
		return n, err == nil
	}
	return parseMatrix(input, "\n", delimiter, parser)
}

func ParseIntArray(input string, delimiter string) []int {
	parser := func(str string) (int, bool) {
		n, err := strconv.Atoi(str)
		return n, err == nil
	}
	return ParseArray(input, delimiter, parser)
}

func ParseCharArray(input string, delimiter string) []string {
	parser := func(str string) (string, bool) { return str, true }
	return ParseArray(input, delimiter, parser)
}

func ParseCharMatrix(input string) [][]string {
	parser := func(str string) (string, bool) { return str, true }
	return parseMatrix(input, "\n", "", parser)
}

func ParseArray[T any](input string, del string, parser func(str string) (T, bool)) []T {
	arr := make([]T, 0)
	cells := strings.Split(input, del)
	for _, cell := range cells {
		num, ok := parser(cell)
		if ok {
			arr = append(arr, num)
		}
	}
	return arr
}

func parseMatrix[T any](input string, lineDel string, cellDel string, parser func(str string) (T, bool)) [][]T {
	data := make([][]T, 0)

	lines := strings.Split(input, lineDel)
	for _, line := range lines {
		data = append(data, ParseArray(line, cellDel, parser))
	}

	return data
}

func ParseNumber(str string, i int) (int, int, bool) {
	if !IsDigit(rune(str[i])) {
		return i + 1, 0, false
	}
	num := 0
	for i < len(str) {
		if IsDigit(rune(str[i])) {
			num = num*10 + int(str[i]) - '0'
		} else {
			return i, num, true
		}
		i++
	}
	return i, num, true
}

func IsDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func Remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func CopySlice[T any](slice []T) []T {
	cp := make([]T, len(slice))
	copy(cp[:], slice)
	return cp
}

func RunSolvers(solutions Solutions) {
	Run(solutions)
}

func Run(solutions Solutions) {
	fmt.Println("\n--- DAY", solutions.Day, "Part 1 ---")
	ok := RunExample(solutions.Day, solutions.Part1)
	if !ok {
		return
	}

	ok = RunReal(solutions.Day, solutions.Part1)
	if !ok {
		return
	}

	fmt.Println("\n--- DAY", solutions.Day, "Part 2 ---")
	ok = RunExample(solutions.Day, solutions.Part2)
	if !ok {
		return
	}

	RunReal(solutions.Day, solutions.Part2)
}

func RunExample(day int, part Solution) bool {
	exampleInput := ReadFile(day, true)
	if part.ExampleInput != "" {
		exampleInput = part.ExampleInput
	}
	example := part.Solver(exampleInput, true)
	pass := example == part.Expected.Example
	if pass {
		fmt.Printf("Example: OK %v\n", example)
	} else {
		fmt.Printf("Example: FAIL %v EXPECTED %v\n", example, part.Expected.Example)
	}
	return pass
}

func RunReal(day int, part Solution) bool {
	start := time.Now()
	input := part.Solver(ReadFile(day, false), false)
	pass := input == part.Expected.Real
	if pass {
		fmt.Printf("Real   : OK %v [%v]\n", input, time.Since(start))
	} else {
		fmt.Printf("Real   : FAIL %v EXPECTED %v [%v]\n", input, part.Expected.Real, time.Since(start))
	}
	return pass
}

type Solutions struct {
	Day   int
	Part1 Solution
	Part2 Solution
}

type Solution struct {
	Solver       Solver
	Expected     Expected
	ExampleInput string
}

type Solver func(input string, example bool) string

type Expected struct {
	Example string
	Real    string
}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

func CountDigits(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func FirstNDigits(num, n int) int {
	return num / int(math.Pow(10, float64(n)))
}

func LastNDigits(num, n int) int {
	return num % int(math.Pow(10, float64(n)))
}

func InBounds(width, height, x, y int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}

func InArr[T any](data [][]T, x, y int) bool {
	return x >= 0 && y >= 0 && y < len(data) && x < len(data[y])
}

func InArrP[T any](data [][]T, p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < len(data) && p.X < len(data[p.Y])
}

func Pop[T any](s []T) ([]T, T) {
	lastIndex := len(s) - 1
	return s[:lastIndex], s[lastIndex]
}

func FloodFill[T any](arr [][]T, start Point, getNeighbors func(arr [][]T, point Point) []Point) *Set[Point] {
	visited := NewSet[Point]()
	queue := []Point{start}
	for len(queue) > 0 {
		var p Point
		queue, p = Pop(queue)
		if visited.Has(p) {
			continue
		}
		visited.Add(p)
		neighbors := getNeighbors(arr, p)
		for _, neighbor := range neighbors {
			queue = append(queue, neighbor)
		}
	}
	return visited
}

func IsInteger(n float64) bool {
	precision := 1e-6
	if _, frac := math.Modf(n); frac < precision || frac > 1-precision {
		return true
	}
	return false
}

func Create2DArray[T any](width, height int) [][]T {
	arr := make([][]T, height)
	for y := range height {
		arr[y] = make([]T, width)
	}
	return arr
}

func PosMod(n, mod int) int {
	return ((n % mod) + mod) % mod
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func PowInt(n, power int) int {
	return int(math.Pow(float64(n), float64(power)))
}

func MapList[T, Y any](arr []T, mapper func(x T) Y) []Y {
	result := make([]Y, len(arr))
	for i, val := range arr {
		result[i] = mapper(val)
	}
	return result
}

func ToString[T any](x T) string {
	return fmt.Sprint(x)
}

func GetCardinalNeighbors(pos Point) []Point {
	neighbors := make([]Point, 0, 4)
	deltas := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, delta := range deltas {
		neighbors = append(neighbors, Point{pos.X + delta.X, pos.Y + delta.Y})
	}
	return neighbors
}

func GetCardinalNeighborsInside(pos Point, xMin, xMax, yMin, yMax int) []Point {
	neighbors := make([]Point, 0, 4)
	for _, n := range GetCardinalNeighbors(pos) {
		if n.Y >= yMin && n.Y <= yMax && n.X >= xMin && n.X <= xMax {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func PointsOfArr[T any](arr [][]T) []Point {
	points := make([]Point, 0, len(arr)*len(arr[0]))
	for y, row := range arr {
		for x := range row {
			points = append(points, Point{x, y})
		}
	}
	return points
}

func PrintMatrix[T any](arr [][]T) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func PrintIntMatrix(arr [][]int, width int) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Printf(fmt.Sprintf("%%%dd", width), cell)
		}
		fmt.Println()
	}
}
