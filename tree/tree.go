package tree

import (
	"sort"
)

type Node struct {
	value     string
	frequency int
	Left      *Node
	Right     *Node
}

func Tree(charsCount map[string]int) Node {
	nodes := nodesList(charsCount)
	return buildTree(nodes)
}

func nodesList(charsCount map[string]int) []Node {
	res := []Node{}
	for ch, cnt := range charsCount {
		n := Node{
			value:     ch,
			frequency: cnt,
			Left:      nil,
			Right:     nil,
		}

		res = append(res, n)
	}

	return res
}

func buildTree(nodes []Node) Node {
	nodes = sortNodes(nodes)

	for len(nodes) > 1 {
		node1 := nodes[0]
		node2 := nodes[1]

		nodes = nodes[2:]

		node := Node{
			value:     "",
			frequency: node1.frequency + node2.frequency,
			Left:      &node1,
			Right:     &node2,
		}

		nodes = append([]Node{node}, nodes...)
		nodes = sortNodes(nodes)
	}

	return nodes[0]
}

func sortNodes(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].frequency < nodes[j].frequency
	})
	return nodes
}
