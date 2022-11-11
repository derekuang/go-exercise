package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var count int

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elemcount: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
	fmt.Println(count)
}

func visit(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		count++
	}
	visit(n.FirstChild)
	visit(n.NextSibling)
}
