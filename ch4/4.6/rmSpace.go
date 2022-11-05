package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []byte("a  b\n\t  c\t def  ")
	fmt.Println(string(rmSpace(s)))
}

func rmSpace(s []byte) []byte {
	p := 0
	for i := 0; i < len(s); p++ {
		if unicode.IsSpace(rune(s[i])) {
			s[p] = ' '
			for i < len(s) && unicode.IsSpace(rune(s[i])) {
				i++
			}
		} else {
			s[p] = s[i]
			i++
		}
	}
	return s[:p]
}
