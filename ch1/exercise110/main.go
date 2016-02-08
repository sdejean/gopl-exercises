// ch1 exercise 1.10 - Fetchall fetches URLs in parallel and reports their
// times and sizes.
// - pulls each url twice
// - saves output to a file for analysis
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	MaxRequests = 2 // max number of test runs per url
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for i := 1; i <= len(os.Args[1:])*MaxRequests; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	for i := 1; i <= MaxRequests; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprint(err) // send to channel ch
			return
		}

		f, err := ioutil.TempFile("", "")
		if err != nil {
			ch <- fmt.Sprintf("while opening a temp file %s: %v", err)
			return
		}
		nbytes, err := io.Copy(f, resp.Body)
		resp.Body.Close() // don't leak resources
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.3fs\t%7d\t%s\t%s", secs, nbytes, url, f.Name())
		f.Close()
	}
}
