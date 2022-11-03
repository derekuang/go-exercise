package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var f384 = flag.Bool("sha384", false, "use SHA384")
var f512 = flag.Bool("sha512", false, "use SHA512")

func main() {
	flag.Parse()
	data, _ := io.ReadAll(bufio.NewReader(os.Stdin))

	if *f384 {
		fmt.Printf("SHA384: %x\n", sha512.Sum384(data))
	} else if *f512 {
		fmt.Printf("SHA512: %x\n", sha512.Sum512(data))
	} else {
		fmt.Printf("SHA256: %x\n", sha256.Sum256(data))
	}
}
