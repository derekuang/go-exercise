package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elemcount: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		visit(n.NextSibling)
		return
	} else if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}

	visit(n.FirstChild)
	visit(n.NextSibling)
}
