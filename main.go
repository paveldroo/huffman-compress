package main

import (
	"log"
	"os"

	"github.com/paveldroo/huffman-compress/codec"
	"github.com/paveldroo/huffman-compress/counter"
	"github.com/paveldroo/huffman-compress/header"
	"github.com/paveldroo/huffman-compress/reader"
	"github.com/paveldroo/huffman-compress/tree"
)

const filenameArgsCount = 2

func main() {
	args := os.Args
	if len(args) < filenameArgsCount {
		log.Fatal("you should specify filename as argument")
	}
	fname := args[1]

	data, err := reader.FileData(fname)
	if err != nil {
		log.Fatal("get filedata: %w", err)
	}

	chars, err := counter.CharsCount(data)
	if err != nil {
		log.Fatalf("can't count chars: %s", err.Error()) //nolint:gosec // no injection
	}

	t := tree.Tree(chars)
	charsTable := codec.CharsCodes(&t)
	h := header.Header(charsTable)
	encodedFileData, err := codec.Encode(h, data, charsTable)
	if err != nil {
		log.Fatalf("encoding failed: %s", err.Error()) //nolint:gosec
	}

	err = os.WriteFile("result", encodedFileData, 0o600) //nolint:gosec,mnd
	if err != nil {
		log.Fatalf("write file: %s", err.Error()) //nolint:gosec
	}
}
