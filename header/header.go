package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"maps"
	"slices"
)

const LenBytes = 4

func Header(charsTable map[string]string) (string, error) {
	b := bytes.Buffer{}

	keys := slices.Sorted(maps.Keys(charsTable))
	for _, key := range keys {
		code := charsTable[key]
		b.WriteString(key + code)
	}

	buf := make([]byte, LenBytes)
	binary.BigEndian.PutUint32(buf, uint32(len(b.String()))) //nolint:gosec // i'm pretty sure
	buf = append(buf, b.Bytes()...)

	res := ""

	for _, c := range buf {
		res += fmt.Sprintf("%08b", c)
	}

	return res, nil
}
