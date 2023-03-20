package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	w := os.Stdout
	f, err := os.Create("result.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	mw := io.MultiWriter(w, f)

	encoder := csv.NewWriter(mw)
	encoder.Write([]string{"a", "b", "c"})
	encoder.WriteAll([][]string{
		{"d", "e", "f"},
		{"g", "h", "i"},
	})
	encoder.Flush()
}
