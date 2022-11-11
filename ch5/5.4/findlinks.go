package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var s string = `<html lang="en" data-theme="auto">
<head>

<link rel="preconnect" href="https://www.googletagmanager.com">
<script >(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>

<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#00add8">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Material+Icons">
<link rel="stylesheet" href="/css/styles.css">


  <script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>

<script src="/js/site.js"></script>
<meta name="og:url" content="https://go.dev/">
<meta name="og:title" content="The Go Programming Language">
<title>The Go Programming Language</title>`

func main() {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			fallthrough
		case "link":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "img":
			fallthrough
		case "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		default:
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}
