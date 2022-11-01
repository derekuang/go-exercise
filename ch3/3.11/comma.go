// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	var i, f string

	dot := strings.LastIndex(s, ".")
	if dot > 0 {
		i, f = s[:dot], s[dot:]
	} else {
		i, f = s, ""
	}

	if i[0] == '+' || i[0] == '-' {
		buf.WriteByte(i[0])
		i = i[1:]
	}

	n := len(i)
	if n <= 3 {
		buf.WriteString(i)
		buf.WriteString(f)
		return buf.String()
	}

	start := n % 3
	buf.WriteString(i[:start])
	for ; start < n; start += 3 {
		if start != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(i[start : start+3])
	}

	buf.WriteString(f)
	return buf.String()
}

//!-
