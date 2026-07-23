package codec

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/paveldroo/huffman-compress/header"
)

func Decode(data []byte) ([]byte, error) {
	headerLen := binary.BigEndian.Uint32(data[:header.LenBytes])

	h := data[header.LenBytes : header.LenBytes+headerLen]
	codesTable := codesTableFromBytes(h)

	fdataBytes := data[header.LenBytes+headerLen:]

	return []byte(decodedData(fdataBytes, codesTable)), nil
}

func codesTableFromBytes(data []byte) map[string]string {
	charsSeq := strings.Split(string(data), ":!")
	res := map[string]string{}
	for i := 0; i < len(charsSeq)-2; i += 2 {
		res[charsSeq[i+1]] = charsSeq[i]
	}

	return res
}

func decodedData(data []byte, codesTable map[string]string) string {
	fdataBinaryStr := strings.Builder{}

	for _, c := range data {
		fmt.Fprintf(&fdataBinaryStr, "%08b", c)
		// fmt.Println(fdataBinaryStr.String())
	}

	b := strings.Builder{}

	cur := ""

	for _, el := range fdataBinaryStr.String() {
		cur += string(el)
		if ch, ok := codesTable[cur]; ok {
			b.WriteString(ch)
			cur = ""
		}
	}

	err := os.WriteFile("../decoded_result", []byte(b.String()), 0o600) //nolint:mnd // permissions code
	if err != nil {
		panic("can't write a decoded file")
	}

	return b.String()
}
