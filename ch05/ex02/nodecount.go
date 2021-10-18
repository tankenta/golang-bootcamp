package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	count := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "nodecount: %v\n", err)
		os.Exit(1)
	}
	nodecount(count, doc)
	for k, v := range count {
		fmt.Println(k, v)
	}
}

func nodecount(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	c := n.FirstChild
	if c != nil {
		nodecount(count, c)
	}
	s := n.NextSibling
	if s != nil {
		nodecount(count, s)
	}
}
