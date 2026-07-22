package header

import (
	"bytes"
	"encoding/binary"
	"maps"
	"slices"
)

const LenBytes = 4

func Header(charsTable map[string]string) []byte {
	b := bytes.Buffer{}

	keys := slices.Sorted(maps.Keys(charsTable))
	for _, key := range keys {
		code := charsTable[key]
		b.WriteString(key + code)
	}

	buf := make([]byte, LenBytes, LenBytes+len(b.Bytes()))
	binary.BigEndian.PutUint32(buf, uint32(len(b.String()))) //nolint:gosec // i'm pretty sure

	return append(buf, b.Bytes()...)
}
