package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urlPrefix := "http://"
	for _, url := range os.Args[1:] {
		urlWithPrefix := ""
		if strings.HasPrefix(url, urlPrefix) {
			urlWithPrefix = url
		} else {
			urlWithPrefix = urlPrefix + url
		}
		resp, err := http.Get(urlWithPrefix)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", urlWithPrefix, err)
			os.Exit(1)
		}
	}
}
