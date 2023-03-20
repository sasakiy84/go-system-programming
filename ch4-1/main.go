package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("seconds is required")
	}

	s, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	d := time.Duration(s) * time.Second

	c := time.After(d)
	<-c

	fmt.Println("It's time")
}
