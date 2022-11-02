package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]
	fmt.Println(isSimilar(s1, s2))
}

func isSimilar(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1, m2 := make(map[rune]int), make(map[rune]int)

	for _, c := range s1 {
		m1[c]++
	}
	for _, c := range s2 {
		m2[c]++
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
