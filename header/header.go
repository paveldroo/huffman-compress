package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"maps"
	"slices"
	"strings"
)

const (
	LenBytes   = 4
	CountBytes = 4
)

func Header(charsTable map[string]string, charCount uint32) (string, error) {
	b := bytes.Buffer{}

	keys := slices.Sorted(maps.Keys(charsTable))
	for _, key := range keys {
		code := charsTable[key]
		b.WriteString(key + ":!" + code + ":!")
	}

	buf := make([]byte, LenBytes+CountBytes, LenBytes+CountBytes+len(b.String()))
	binary.BigEndian.PutUint32(buf[:LenBytes], uint32(len(b.String()))) //nolint:gosec // i'm pretty sure
	binary.BigEndian.PutUint32(buf[LenBytes:LenBytes+CountBytes], charCount)
	buf = append(buf, b.Bytes()...)

	res := strings.Builder{}

	for _, c := range buf {
		fmt.Fprintf(&res, "%08b", c)
	}

	return res.String(), nil
}
