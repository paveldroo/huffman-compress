package codec_test

import (
	"maps"
	"testing"

	"github.com/paveldroo/huffman-compress/codec"
	"github.com/paveldroo/huffman-compress/tree"
)

func TestCharsCodes(t *testing.T) {
	t.Parallel()
	charsCount := map[string]int{
		"C": 32,
		"D": 42,
		"E": 120,
		"K": 7,
		"L": 42,
		"M": 24,
		"U": 37,
		"Z": 2,
	}
	want := map[string]string{
		"C": "1110",
		"D": "101",
		"E": "0",
		"K": "111101",
		"L": "110",
		"M": "11111",
		"U": "100",
		"Z": "111100",
	}

	root := tree.Tree(charsCount)

	chCodes := codec.CharsCodes(&root)
	if !maps.Equal(chCodes, want) {
		t.Fatalf("not equal\nresult: %v\nwant: %v", chCodes, want)
	}
}
