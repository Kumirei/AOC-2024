package d23

import (
	"awesomeProject/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solutions() util.Solutions {
	return util.Solutions{
		Day: 23,
		Part1: util.Solution{
			Solver:   Part1,
			Expected: util.Expected{Example: "7", Real: "1064"},
		},
		Part2: util.Solution{
			Solver:   Part2,
			Expected: util.Expected{Example: "co,de,ka,ta", Real: "aq,cc,ea,gc,jo,od,pa,rg,rv,ub,ul,vr,yy"},
		},
	}
}

func Part1(input string, example bool) string {
	data := parse(input)
	//fmt.Println(data)

	nodes := getNodes(data)

	//fmt.Println(*nodes["tc"].edges)
	fullyConnected := make([][]*Node, 0)
	for _, start := range nodes {
		perms := permutations(start, 3)
		for _, perm := range perms {
			connected := isFullyConnected(perm)
			if connected {
				fullyConnected = append(fullyConnected, perm)
			}
		}
	}

	hasStartWithT := make([][]*Node, 0)
	for _, perm := range fullyConnected {
		for _, n := range perm {
			if n.name[0] == 't' {
				hasStartWithT = append(hasStartWithT, perm)
				break
			}
		}
	}

	//fmt.Println("PERM", len(hasStartWithT)/3)

	return strconv.Itoa(len(hasStartWithT) / 3)
}

func Part2(input string, example bool) string {
	connections := parsePart2(input)

	largest := util.NewSet[string]()
	for node := range connections {
		clique := dfs(connections, node, util.NewSet[string]())
		if clique.Size() > largest.Size() {
			largest = clique
		}
	}

	list := largest.ToList()
	sort.Strings(list)
	answer := strings.Join(list, ",")

	return answer
}

func dfs(connections map[string]*util.Set[string], node string, visited *util.Set[string]) *util.Set[string] {
	visited.Add(node)

	largestParty := util.NewSet[string]()
outer:
	for connection := range connections[node].List {
		if visited.Has(connection) {
			continue
		}
		for prev := range visited.List {
			if !connections[connection].Has(prev) {
				continue outer
			}
		}
		party := dfs(connections, connection, visited)
		if party.Size() > largestParty.Size() {
			largestParty = party
		}
	}
	if largestParty.Size() == 0 {
		largestParty = visited.Clone()
	}
	return largestParty
}

func (n *Node) getNeighborSet() *util.Set[*Node] {
	neighbors := n.neighbors()
	neighborSet := util.NewSet[*Node]()
	neighborSet.AddMulti(neighbors...)
	return neighborSet
}

func getNodes(data [][]string) map[string]*Node {
	nodes := make(map[string]*Node)
	for _, pair := range data {
		keyA := pair[0]
		keyB := pair[1]
		a := nodes[keyA]
		b := nodes[keyB]
		if a == nil {
			a = &Node{name: keyA, edges: util.NewSet[Edge]()}
		}
		if b == nil {
			b = &Node{name: keyB, edges: util.NewSet[Edge]()}
		}
		edge := Edge{a, b}
		a.edges.Add(edge)
		b.edges.Add(edge)
		nodes[keyA] = a
		nodes[keyB] = b
	}
	return nodes
}

func isFullyConnected(nodes []*Node) bool {
	for i := range len(nodes) {
		for j := i + 1; j < len(nodes); j++ {
			ok := nodes[i].hasNeighbor(nodes[j])
			if !ok {
				return false
			}
		}
	}
	return true
}

func (n *Node) hasNeighbor(neighbor *Node) bool {
	for edge := range n.edges.List {
		if edge.a == neighbor || edge.b == neighbor {
			return true
		}
	}
	return false
}

func permutations(n *Node, size int) [][]*Node {
	perms := make([][]*Node, 0)
	neighbors := n.neighbors()
	for i := range len(neighbors) {
		for j := i + 1; j < len(neighbors); j++ {
			perms = append(perms, []*Node{n, neighbors[i], neighbors[j]})
		}
	}
	return perms
}

func (n *Node) String() string {
	return n.name
}

func (e *Edge) String() string {
	return fmt.Sprintf("%s-%s", e.a, e.b)
}

func (n *Node) neighbors() []*Node {
	arr := make([]*Node, n.edges.Size())
	i := 0
	for edge := range n.edges.List {
		var next *Node
		if edge.a == n {
			next = edge.b
		} else {
			next = edge.a
		}
		arr[i] = next
		i++
	}
	return arr
}

type Node struct {
	name  string
	edges *util.Set[Edge]
}

type Edge struct {
	a *Node
	b *Node
}

func parse(input string) [][]string {
	return util.ParseStringMatrix(input, "-")
}

func parsePart2(input string) map[string]*util.Set[string] {
	parsed := make(map[string]*util.Set[string])
	data := util.ParseStringMatrix(input, "-")
	for _, pair := range data {
		a, b := pair[0], pair[1]
		if parsed[a] == nil {
			parsed[a] = util.NewSet[string]()
		}
		if parsed[b] == nil {
			parsed[b] = util.NewSet[string]()
		}
		parsed[a].Add(b)
		parsed[b].Add(a)
	}
	return parsed
}
