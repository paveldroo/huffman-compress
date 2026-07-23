package codec

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/paveldroo/huffman-compress/tree"
)

const bitsCount = 8

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
	b := bytes.Buffer{}
	for i := 0; i < len(bitStr)-1; i += bitsCount {
		right := min(i+bitsCount, len(bitStr))

		chunk := bitStr[i:right]
		if len(chunk) < bitsCount {
			chunk += strings.Repeat("0", bitsCount-len(chunk))
		}
		v, err := strconv.ParseUint(chunk, 2, bitsCount)
		if err != nil {
			return nil, fmt.Errorf("convert bit string to byte: %w", err)
		}
		b.WriteByte(byte(v))
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
