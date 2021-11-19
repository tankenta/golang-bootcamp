package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dumptextnode: %v\n", err)
		os.Exit(1)
	}
	for _, nodetext := range visit(nil, doc) {
		fmt.Println(nodetext)
	}
}

func visit(nodetexts []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		nodetexts = append(nodetexts, n.Data)
	}
	c := n.FirstChild
	if c != nil {
		nodetexts = visit(nodetexts, c)
	}
	s := n.NextSibling
	if s != nil {
		nodetexts = visit(nodetexts, s)
	}
	return nodetexts
}
