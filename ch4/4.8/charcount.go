package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters types

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsDigit(r) {
			counts["digit"]++
		} else {
			counts["other"]++
		}
	}
	fmt.Printf("type\tcount\n")
	for t, n := range counts {
		fmt.Printf("%s\t%d\n", t, n)
	}
}
