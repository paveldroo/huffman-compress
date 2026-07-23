package codec

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/paveldroo/huffman-compress/header"
)

const headerStart = header.LenBytes + header.CountBytes

func Decode(data []byte) ([]byte, error) {
	headerLen := binary.BigEndian.Uint32(data[:header.LenBytes])
	charCount := binary.BigEndian.Uint32(data[header.LenBytes:headerStart])

	h := data[headerStart : headerStart+headerLen]
	codesTable := codesTableFromBytes(h)

	fdataBytes := data[headerStart+headerLen:]

	return []byte(decodedData(fdataBytes, codesTable, charCount)), nil
}

func codesTableFromBytes(data []byte) map[string]string {
	charsSeq := strings.Split(string(data), ":!")
	res := map[string]string{}
	for i := 0; i < len(charsSeq)-2; i += 2 {
		res[charsSeq[i+1]] = charsSeq[i]
	}

	return res
}

func decodedData(data []byte, codesTable map[string]string, charCount uint32) string {
	bits := strings.Builder{}
	for _, by := range data {
		fmt.Fprintf(&bits, "%08b", by)
	}

	b := strings.Builder{}
	cur := strings.Builder{}

	var emitted uint32
	for _, bit := range bits.String() {
		cur.WriteRune(bit)

		if ch, ok := codesTable[cur.String()]; ok {
			b.WriteString(ch)
			cur.Reset()

			emitted++
			if emitted == charCount {
				return b.String()
			}
		}
	}

	return b.String()
}
