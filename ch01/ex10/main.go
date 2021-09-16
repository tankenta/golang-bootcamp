package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// 第1引数を取得ページ出力ファイルとする
	f, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	ch := make(chan string)
	// 第2引数以降がURL．ただし複数指定しても出力ファイルには最後の1ページしか記録されない
	for _, url := range os.Args[2:] {
		go fetch(f, url, ch)
	}
	for range os.Args[2:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	f.Close()
}

func fetch(f *os.File, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
