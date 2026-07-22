package header_test

import (
	"encoding/binary"
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

	str := "a100001b0c10d01"

	buf := make([]byte, header.LenBytes)
	binary.BigEndian.PutUint32(buf, uint32(len(str))) //nolint:gosec // i'm pretty sure
	want := string(buf) + str
	h, err := header.Header(charsTable)
	if err != nil {
		t.Fatalf("compose header: %s", err.Error())
	}

	if want != h {
		t.Fatalf("not equal, want: %s, header: %s\n", want, h)
	}
}
