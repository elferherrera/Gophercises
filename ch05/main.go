package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/elferherrera/Gophercises/ch04/link"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url for the map")
	flag.Parse()

	fmt.Printf("Collecing information from: %v\n", *urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	fmt.Println(links)

}

/*
	1.- GET the webpage
	2.- parse all the links on the page
	3.- build proper url with our links
	4.- filter out any links with a diff domain
	5.- find all the pages (BFS)
	6.- print out XML
*/
