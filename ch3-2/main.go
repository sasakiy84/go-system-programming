package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := rand.Reader
	io.CopyN(f, r, int64(1024))

	// use `wc --bytes temp.txt` command for check
}
