package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	for range s {
		rotate(s)
		fmt.Println(s)
	}
}

// rotate rotates a slice in clockwise
func rotate(s []int) {
	slen := len(s)
	if slen < 1 {
		return
	}

	last := s[slen-1]
	for i := slen - 1; i > 0; i-- {
		s[i] = s[i-1]
	}
	s[0] = last
}
