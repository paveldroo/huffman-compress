package header_test

import (
	"encoding/binary"
	"slices"
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
	flat := "a100001b0c10d01"
	want := make([]byte, header.LenBytes, header.LenBytes+len(flat))
	binary.BigEndian.PutUint32(want, uint32(len(flat))) //nolint:gosec // i'm pretty sure
	want = append(want, []byte(flat)...)
	h, err := header.Header(charsTable)
	if err != nil {
		t.Fatalf("compose header: %s", err.Error())
	}

	if !slices.Equal(want, h) {
		t.Fatalf("not equal, want: %v, header: %v\n", want, h)
	}
}
