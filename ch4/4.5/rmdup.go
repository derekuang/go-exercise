package main

import "fmt"

func main() {
	s := []string{"abc", "def", "def", "abc", "123", "123"}
	fmt.Println(rmDup(s))
}

func rmDup(s []string) []string {
	p := 0
	for i, str := range s {
		if i == 0 {
			continue
		}
		if str != s[p] {
			s[p+1] = str
			p++
		}
	}
	return s[:p+1]
}
