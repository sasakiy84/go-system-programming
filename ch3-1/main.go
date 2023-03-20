package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// parse arguments
	if len(os.Args) != 3 {
		log.Fatal("expected 2 arguments")
	}
	src, dst := os.Args[1], os.Args[2]

	// process copy
	fr, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()

	fw, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()
	io.Copy(fw, fr)
}
