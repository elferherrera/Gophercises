package main

import (
	"fjherrera/ch04/link"
	"flag"
	"fmt"
	"os"
)

func main() {
	htmlFile := flag.String("page", "pages/ex1.html", "html file name")
	flag.Parse()

	f, err := os.Open(*htmlFile)
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
