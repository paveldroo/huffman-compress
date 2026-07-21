package tree

import (
	"slices"
)

type Node struct {
	Value     string
	Frequency int
	Left      *Node
	Right     *Node
}

func Tree(charsCount map[string]int) Node {
	nodes := nodesList(charsCount)

	return buildTree(nodes)
}

func nodesList(charsCount map[string]int) []Node {
	res := make([]Node, 0, len(charsCount))
	for ch, cnt := range charsCount {
		n := Node{
			Value:     ch,
			Frequency: cnt,
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
			Value:     "",
			Frequency: node1.Frequency + node2.Frequency,
			Left:      &node1,
			Right:     &node2,
		}

		nodes = append([]Node{node}, nodes...)
		nodes = sortNodes(nodes)
	}

	return nodes[0]
}

func sortNodes(nodes []Node) []Node {
	slices.SortFunc(nodes, func(a, b Node) int {
		if a.Frequency == b.Frequency && a.Value < b.Value {
			return -1
		}

		if a.Frequency < b.Frequency {
			return -1
		}

		return 1
	})

	return nodes
}
