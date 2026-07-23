package header_test

import (
	"encoding/binary"
	"fmt"
	"strings"
	"testing"

	"github.com/paveldroo/huffman-compress/header"
)

func TestHeader(t *testing.T) {
	t.Parallel()
	charsTable := map[string]string{
		"a": "100001",
		"b": "0",
		"c": "10",
		"d": "01",
	}

	const charCount = 0

	// sorted keys joined as key:!code:!
	body := "a:!100001:!b:!0:!c:!10:!d:!01:!"

	buf := make([]byte, header.LenBytes+header.CountBytes)
	binary.BigEndian.PutUint32(buf[:header.LenBytes], uint32(len(body))) //nolint:gosec // test data
	binary.BigEndian.PutUint32(buf[header.LenBytes:], charCount)
	buf = append(buf, body...)

	frame := strings.Builder{}
	for _, c := range buf {
		fmt.Fprintf(&frame, "%08b", c)
	}
	want := frame.String()

	h, err := header.Header(charsTable, charCount)
	if err != nil {
		t.Fatalf("compose header: %s", err.Error())
	}

	if want != h {
		t.Fatalf("not equal, want: %s, header: %s\n", want, h)
	}
}
