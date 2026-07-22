package codec

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/paveldroo/huffman-compress/tree"
)

var (
	errNoCode = errors.New("no code for char")
	charsMap  = make(map[string]string) //nolint:gochecknoglobals // need global for algo
)

func Encode(header string, data []byte, charsTable map[string]string) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(header)
	for _, char := range string(data) {
		code, ok := charsTable[string(char)]
		if !ok {
			return nil, fmt.Errorf("%w: %s", errNoCode, string(char))
		}

		buf.WriteString(code)
	}

	res, err := ConvertToBytes(buf.String())
	if err != nil {
		return nil, fmt.Errorf("convert all data to bytes: %w", err)
	}

	return res, nil
}

func ConvertToBytes(bitStr string) ([]byte, error) {
	bitsLen := (len(bitStr) + 7) / 8

	b := bytes.Buffer{}
	lastIdx := 0
	for range bitsLen {
		right := lastIdx + 8
		if right >= len(bitStr) {
			right = len(bitStr) - 1
		}

		if lastIdx == right {
			continue
		}

		oneBitStr := bitStr[lastIdx:right]
		byteToAdd, err := strconv.ParseUint(oneBitStr, 2, 8)
		if err != nil {
			return nil, fmt.Errorf("convert bit string to byte: %w", err)
		}
		b.WriteByte(byte(byteToAdd))
		lastIdx = lastIdx + 8
	}

	return b.Bytes(), nil
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
