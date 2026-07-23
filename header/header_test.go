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

	str := "000000000000000000000000000111110110000100111010001000010011000100110000001100000" +
		"0110000001100000011000100111010001000010110001000111010001000010011000000111010001000010" +
		"1100011001110100010000100110001001100000011101000100001011001000011101000100001001100000" +
		"01100010011101000100001"

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
