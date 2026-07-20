package main

import (
	"fmt"
	"log"
	"os"

	"github.com/paveldroo/huffman-compress/counter"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("you should specify filename as argument")
		os.Exit(1)
	}
	fname := args[1]

	chars, err := counter.CharsCount(fname)
	if err != nil {
		log.Fatalf("can't count chars: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(chars)
}
