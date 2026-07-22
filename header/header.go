package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"maps"
	"slices"

	"github.com/paveldroo/huffman-compress/codec"
)

const LenBytes = 4

func Header(charsTable map[string]string) ([]byte, error) {
	b := bytes.Buffer{}

	keys := slices.Sorted(maps.Keys(charsTable))
	for _, key := range keys {
		code := charsTable[key]
		byteCode, err := codec.ConvertToBytes(code)
		if err != nil {
			return nil, fmt.Errorf("convert code to bytes: %w", err)
		}
		b.WriteString(key)
		b.Write(byteCode)
	}

	buf := make([]byte, LenBytes, LenBytes+len(b.Bytes()))
	binary.BigEndian.PutUint32(buf, uint32(len(b.String()))) //nolint:gosec // i'm pretty sure

	return append(buf, b.Bytes()...), nil
}
