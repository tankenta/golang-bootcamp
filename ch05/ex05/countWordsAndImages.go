package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Errorf("%s", err)
		}
		fmt.Println("words: ", words)
		fmt.Println("images: ", images)
		fmt.Println()
	}
}

func visitImage(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	c := n.FirstChild
	if c != nil {
		links = visitImage(links, c)
	}
	s := n.NextSibling
	if s != nil {
		links = visitImage(links, s)
	}
	return links
}

func visitNodeText(nodetexts []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		nodetexts = append(nodetexts, n.Data)
	}
	c := n.FirstChild
	if c != nil {
		nodetexts = visitNodeText(nodetexts, c)
	}
	s := n.NextSibling
	if s != nil {
		nodetexts = visitNodeText(nodetexts, s)
	}
	return nodetexts
}

func countWords(nodetexts []string) int {
	wholeText := strings.Join(nodetexts, " ")
	input := bufio.NewScanner(strings.NewReader(wholeText))
	input.Split(bufio.ScanWords)
	var counts int
	for input.Scan() {
		counts++
	}
	return counts
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	nodetexts := visitNodeText(nil, doc)
	words = countWords(nodetexts)
	imglinks := visitImage(nil, doc)
	images = len(imglinks)
	return words, images
}
