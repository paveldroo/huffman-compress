package main

import (
	"log"
	"os"

	"github.com/paveldroo/huffman-compress/counter"
	"github.com/paveldroo/huffman-compress/tree"
)

const filenameArgsCount = 2

func main() {
	args := os.Args
	if len(args) < filenameArgsCount {
		log.Fatal("you should specify filename as argument")
	}
	fname := args[1]

	chars, err := counter.CharsCount(fname)
	if err != nil {
		log.Fatalf("can't count chars: %s", err.Error()) //nolint:gosec // no injection
	}

	_ = tree.Tree(chars)

	// fmt.Println(tree)
}
