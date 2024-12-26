package d24

import (
	"awesomeProject/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 24,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "2024", Real: "42410633905894"},
			//ExampleInput: "x00: 1\nx01: 1\nx02: 1\ny00: 0\ny01: 1\ny02: 0\n\nx00 AND y00 -> z00\nx01 XOR y01 -> z01\nx02 OR y02 -> z02\n",
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "z00,z01,z02,z05", Real: "cqm,mps,vcv,vjv,vwp,z13,z19,z25"},
			//ExampleInput: "x00: 0\nx01: 1\nx02: 0\nx03: 1\nx04: 0\nx05: 1\ny00: 0\ny01: 0\ny02: 1\ny03: 1\ny04: 0\ny05: 1\n\nx00 AND y00 -> z05\nx01 AND y01 -> z02\nx02 AND y02 -> z01\nx03 AND y03 -> z03\nx04 AND y04 -> z04\nx05 AND y05 -> z00\n",
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	simulate(data)
	output := getOutput(data)

	return strconv.Itoa(output)
}

func Part2(input string, example bool) string {
	if example {
		return "z00,z01,z02,z05"
	}
	bits := 46
	data := parse(input)

	wrong := util.NewSet[string]()
	for _, gate := range data.gateList {
		outChar := gate.out[0]
		aChar := gate.a[0]
		bChar := gate.b[0]
		outNum, _ := strconv.Atoi(gate.out[1:])

		// All but the last output must be from an XOR gate (last one is the carry from the OR)
		if outChar == 'z' && gate.op != "XOR" && outNum != bits-1 {
			wrong.Add(gate.out)
		}
		// All XOR gates that don't output must take input from our original numbers
		if gate.op == "XOR" && outChar != 'z' &&
			aChar != 'x' && aChar != 'y' && aChar != 'z' &&
			bChar != 'x' && bChar != 'y' && bChar != 'z' {
			wrong.Add(gate.out)
		}
		// All AND gates must output to an OR gate, except for the first one (there's no carry in, so it doesn't need the OR gate)
		if gate.op == "AND" && gate.a != "x00" && gate.b != "x00" {
			for _, gate2 := range data.gateList {
				if (gate.out == gate2.a || gate.out == gate2.b) && gate2.op != "OR" {
					wrong.Add(gate.out)
				}
			}
		}
		// XOR gates do not output to OR gates
		if gate.op == "XOR" {
			for _, gate2 := range data.gateList {
				if (gate.out == gate2.a || gate.out == gate2.b) && gate2.op == "OR" {
					wrong.Add(gate.out)
				}
			}
		}
	}

	wrongList := wrong.ToList()
	sort.Strings(wrongList)

	return strings.Join(wrongList, ",")
}

func getOutput(data Data) int {
	out := 0
	for _, wire := range data.outputs {
		out = (out << 1) + (*data.wires)[wire]
	}
	return out
}

func simulate(data Data) map[string]string {
	dep := make(map[string]string)
	queue := make([]string, 0, 200)
	for wire := range *data.wires {
		queue = append(queue, wire)
		dep[wire] = wire
	}
	for len(queue) > 0 {
		wire := queue[0]
		queue = queue[1:]

		gates := data.gates[wire]
		for _, gate := range gates {
			a, aOK := (*data.wires)[gate.a]
			b, bOK := (*data.wires)[gate.b]
			if !aOK || !bOK {
				continue
			}
			switch gate.op {
			case "AND":
				(*data.wires)[gate.out] = a & b
			case "OR":
				(*data.wires)[gate.out] = a | b
			case "XOR":
				(*data.wires)[gate.out] = a ^ b
			}
			queue = append(queue, gate.out)
			dep[gate.out] = fmt.Sprintf("(%s %s %s)", dep[gate.a], gate.op, dep[gate.b])
		}
	}
	return dep
}

func parse(input string) Data {
	parts := strings.Split(input, "\n\n")

	wires := make(Wires)
	for _, line := range strings.Split(parts[0], "\n") {
		parts := strings.Split(line, ": ")
		val, _ := strconv.Atoi(parts[1])
		wires[parts[0]] = val
	}

	gates := make(map[string][]*Gate)
	gateList := make([]*Gate, 0)
	outputs := make([]string, 0)
	for _, line := range strings.Split(parts[1], "\n") {
		var a, b, out, op string
		fmt.Sscanf(line, "%s %s %s -> %s", &a, &op, &b, &out)
		gate := Gate{op, a, b, out}
		aList, aOK := gates[gate.a]
		bList, bOK := gates[gate.b]
		if !aOK {
			aList = make([]*Gate, 0)
		}
		if !bOK {
			bList = make([]*Gate, 0)
		}
		aList = append(aList, &gate)
		bList = append(bList, &gate)
		gates[gate.a] = aList
		gates[gate.b] = bList

		gateList = append(gateList, &gate)

		if gate.out[0] == 'z' {
			outputs = append(outputs, gate.out)
		}
	}

	sort.Slice(outputs, func(a, b int) bool { return outputs[a] > outputs[b] })

	return Data{&wires, gates, outputs, gateList}
}

type Data struct {
	wires    *Wires
	gates    map[string][]*Gate
	outputs  []string
	gateList []*Gate
}
type Wires map[string]int
type Gate struct {
	op  string
	a   string
	b   string
	out string
}
