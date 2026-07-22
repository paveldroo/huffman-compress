package codec

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/paveldroo/huffman-compress/tree"
)

var (
	errNoCode = errors.New("no code for char")
	charsMap  = make(map[string]string) //nolint:gochecknoglobals // need global for algo
)

func Encode(header []byte, data []byte, charsTable map[string]string) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.Write(header)
	for _, char := range string(data) {
		code, ok := charsTable[string(char)]
		if !ok {
			return nil, fmt.Errorf("%w: %s", errNoCode, string(char))
		}
		buf.WriteString(code)
	}

	return buf.Bytes(), nil
}

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
