package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer := 1

	for _, ch := range input {

		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}

		// Capital letters are between 65 to 90
		//if ch >= 'A' && ch <= 'Z' {
		//	fmt.Printf("Char: %#U \n", ch)
		//	answer++
		//}
	}

	fmt.Println("Result: ", answer)
}
