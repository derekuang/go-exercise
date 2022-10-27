// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"go-exercise/ch2/popcount"
)

// -- Alternative implementations --

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

// exercise 2.5
func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

// exercise 2.4
func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

// GO 1.19
// goos: linux
// goarch: amd64
// pkg: go-exercise/ch2/popcount
// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkPopCount
// BenchmarkPopCount-2                     1000000000               0.6571 ns/op          0 B/op          0 allocs/op
// BenchmarkPopCountByLoop
// BenchmarkPopCountByLoop-2               224322742                5.119 ns/op           0 B/op          0 allocs/op
// BenchmarkBitCount
// BenchmarkBitCount-2                     1000000000               0.4275 ns/op          0 B/op          0 allocs/op
// BenchmarkPopCountByClearing
// BenchmarkPopCountByClearing-2           52872973                23.62 ns/op            0 B/op          0 allocs/op
// BenchmarkPopCountByShifting
// BenchmarkPopCountByShifting-2			41270934                27.22 ns/op            0 B/op          0 allocs/op
