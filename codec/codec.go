package codec

import (
	"github.com/paveldroo/huffman-compress/tree"
)

var charsMap = make(map[string]string) //nolint:gochecknoglobals // need global for algo

func CharsCodes(node *tree.Node) map[string]string {
	traverseTree(node, "")

	return charsMap
}

func traverseTree(node *tree.Node, curCode string) {
	left := node.Left
	right := node.Right
	if left != nil {
		curCode += "0"
		if left.Value == "" {
			traverseTree(left, curCode)
		} else {
			charsMap[left.Value] = curCode
		}
		curCode = curCode[:len(curCode)-1]
	}
	if right != nil {
		curCode += "1"
		if right.Value == "" {
			traverseTree(right, curCode)
		} else {
			charsMap[right.Value] = curCode
		}
	}
}
