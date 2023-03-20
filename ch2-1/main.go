package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// use OpenFile but Open or Create, because we append text
	f, err := os.OpenFile("result.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintf(f, "%d, %f %s\n", 1234, 0.3, "aaa")
	if err != nil {
		log.Fatal(err)
	}
}
