package echo_test

import (
	"strings"
	"testing"
)

var s []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

// -- Alternative implementations --

func StringJoinLoop(arr []string) string {
	s, sep := "", ""
	for _, c := range arr {
		s += sep + c
		sep = " "
	}
	return s
}

func StringJoinBuiltin(arr []string) string {
	return strings.Join(arr, " ")
}

// -- Benchmarks --

func BenchmarkStringJoinLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoinLoop(s)
	}
}

func BenchmarkStringJoinBuiltin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoinBuiltin(s)
	}
}

// Go 1.19, Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// goos: linux
// goarch: amd64
// pkg: go-exercise/ch1/1.3
// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkStringJoinLoop
// BenchmarkStringJoinLoop-2        2039856               496.5 ns/op           128 B/op          9 allocs/op
// BenchmarkStringJoinBuiltin
// BenchmarkStringJoinBuiltin-2     8699200               140.5 ns/op            24 B/op          1 allocs/op
