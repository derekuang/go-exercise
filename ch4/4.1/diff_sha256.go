package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(diff_sha256(c1, c2))
}

func diff_sha256(c1, c2 [32]byte) int {
	c := 0
	for i, _ := range [32]byte{} {
		for j := 0; j < 8; j++ {
			if ((c1[i] >> j) & 1) != ((c2[i] >> j) & 1) {
				c++
			}
		}
	}
	return c
}
