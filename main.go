package main

import (
	"log"
	"os"
	"unicode/utf8"

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
		log.Fatalf("get filedata: %s", err.Error()) //nolint:gosec // no injection
	}

	chars := counter.CharsCount(data)

	t := tree.Tree(chars)
	charsTable := codec.CharsCodes(&t)
	charCount := uint32(utf8.RuneCount(data)) //nolint:gosec // fits file size
	h := header.Header(charsTable, charCount)
	encodedFileData, err := codec.Encode(h, data, charsTable)
	if err != nil {
		log.Fatalf("encoding failed: %s", err.Error()) //nolint:gosec
	}

	err = os.WriteFile("encoded_result", encodedFileData, 0o600) //nolint:gosec,mnd
	if err != nil {
		log.Fatalf("write file: %s", err.Error()) //nolint:gosec
	}

	data, err = os.ReadFile("encoded_result")
	if err != nil {
		log.Fatalf("read encoded_result file: %s", err.Error())
	}

	d, err := codec.Decode(data)
	if err != nil {
		log.Fatalf("decode byte data to text: %s", err.Error())
	}

	err = os.WriteFile("decoded_result", d, 0o600) //nolint:gosec,mnd
	if err != nil {
		log.Fatalf("write decoded_result file: %s", err.Error())
	}
	log.Println("decoded_result successfully created!")
}
