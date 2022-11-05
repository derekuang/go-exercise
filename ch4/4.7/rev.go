package main

import "fmt"

func main() {
	s := "Hello, 世界"
	s = string(reverse([]byte(s)))
	fmt.Println(s)
}

func reverse(s []byte) []byte {
	t := []byte{}
	for i, j := len(s)-1, len(s); i >= 0; i, j = i-1, i {
		for s[i]>>6 == 0x2 {
			i--
		}
		t = append(t, s[i:j]...)
	}
	return t
}
