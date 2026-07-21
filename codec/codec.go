package codec

import (
	"fmt"

	"github.com/paveldroo/huffman-compress/tree"
)

var charsMap = make(map[string]string)

func CharsCodes(node *tree.Node) map[string]string {
	traverseTree(node, "")
	return charsMap
}

func traverseTree(node *tree.Node, curCode string) {
	left := node.Left
	right := node.Right
	if left != nil {
		curCode += "0"
		fmt.Printf("before: left freq: %d, left value: %s, curCode: %s\n", left.Frequency, left.Value, curCode)
		if left.Value == "" {
			traverseTree(left, curCode)
		} else {
			charsMap[left.Value] = curCode
		}
		curCode = curCode[:len(curCode)-1]
		fmt.Printf("after: left freq: %d, left value: %s, curCode: %s\n", left.Frequency, left.Value, curCode)
	}
	if right != nil {
		curCode += "1"
		fmt.Printf("before: right freq: %d, right value: %s, curCode: %s\n", right.Frequency, right.Value, curCode)
		if right.Value == "" {
			traverseTree(right, curCode)
		} else {
			charsMap[right.Value] = curCode
		}
		curCode = curCode[:len(curCode)-1]
		fmt.Printf("after: right freq: %d, right value: %s, curCode: %s\n", right.Frequency, right.Value, curCode)
	}
}
