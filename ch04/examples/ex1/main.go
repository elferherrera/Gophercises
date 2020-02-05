package main

import (
	"fjherrera/ch04/link"
	"fmt"
	"strings"
)

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
  <a href="/other-1">
	  how cute <span>this is<span>
	  A link to another page
  </a>
  <a href="/other-2">A third link to a page </a>
</body>
</html>
`

func main() {
	// Example reading strings with a reader instead of []bite
	// Readir for strings
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
