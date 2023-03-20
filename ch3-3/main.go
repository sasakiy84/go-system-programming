package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

var content string = "aaa"

func main() {
	f, err := os.Create("new.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	zipWriter := zip.NewWriter(f)
	defer zipWriter.Close()

	w, err := zipWriter.Create("newfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(content)
	io.Copy(w, r)
}
