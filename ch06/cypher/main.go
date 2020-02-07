package main

import (
	"fmt"
	"strings"
)

func main() {
	var lenght, delta int
	var input string

	fmt.Scanf("%d\n", &lenght)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%delta\n", &delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := strings.ToUpper(alphabetLower)

	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}
	}

	fmt.Println(ret)
}

func rotate(s rune, delta int, key []rune) rune {

	idx := strings.IndexRune(string(key), s)

	//idx := -1
	//for i, r := range key {
	//	if r == s {
	//		idx = i
	//		break
	//	}
	//}

	if idx < 0 {
		panic("idx < 0")
	}

	idx = (idx + delta) % len(key)

	//for i := 0; i < delta; i++ {
	//	idx++
	//	if idx >= len(key) {
	//		idx = 0
	//	}
	//}

	return key[idx]
}
