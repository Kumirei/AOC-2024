package d9

import (
	"awesomeProject/util"
	"strconv"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 9,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "1928", Real: "6154342787400"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "2858", Real: "6183632723350"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)

	left, right := 0, len(data)-1
	for left < right {
		if data[left].Type == Empty && data[right].Type == File {
			temp := data[left]
			data[left] = data[right]
			data[right] = temp
		} else if data[left].Type == File {
			left++
		} else if data[right].Type == Empty {
			right--
		}
	}

	sum := 0
	for i, block := range data {
		if block.Type == Empty {
			break
		}
		sum += block.Id * i
	}

	return strconv.Itoa(sum)
}

func Part2(input string, example bool) string {
	data := parse(input)

	for right := len(data) - 1; right > 0; right-- {
		if data[right].Type == File && !sameFile(data[right], data[right-1]) {
			dataSize := fileSize(data, right)
			for left := 0; left < right; left++ {
				if data[left].Type == Empty && (left == 0 || data[left-1].Type == File) {
					if fileSize(data, left) >= dataSize {
						file := data[right]
						right2 := right
						for right2 < len(data) && sameFile(file, data[right2]) {
							temp := data[left]
							data[left] = data[right2]
							data[right2] = temp
							right2++
							left++
						}

						break
					}
				}
			}
		}
	}

	sum := 0
	for i, block := range data {
		sum += i * block.Id
	}

	return strconv.Itoa(sum)
}

func sameFile(a, b Block) bool {
	return a.Type == b.Type && a.Id == b.Id
}

func fileSize(data Data, i int) (count int) {
	id := data[i].Id
	t := data[i].Type
	for i < len(data) && data[i].Type == t && data[i].Id == id {
		count++
		i++
	}
	return
}

func parse(input string) Data {
	empty := false
	i := 0

	blocks := make(Data, 0)
	for _, char := range input {
		for range char - '0' {
			if empty {
				blocks = append(blocks, Block{Empty, 0})
			} else {
				blocks = append(blocks, Block{File, i})
			}
		}
		empty = !empty
		if !empty {
			i++
		}
	}

	return blocks
}

type Blocks struct {
	Type   BlockType
	Length int
	Id     int
}

type Data []Block

type Block struct {
	Type BlockType
	Id   int
}

type BlockType int

const (
	File BlockType = iota
	Empty
)
