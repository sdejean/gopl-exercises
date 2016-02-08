// gopl/ch1/fetch-mod - combination of exercises 1.7-1.9
// - use io.Copy() instead of ioutil.ReadAll()
// - prepend protocol string 'http://' if user leaves it out
// - print HTTP request status
// - extra:
//   - print request protocol/version info (ex. like curl -i)
//   - print request location (ex. like curl -i)
//   - print content-length in bytes (ex. like curl -i)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// check for protocol prefix
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "%s %s\n", resp.Proto, resp.Status)
		fmt.Fprintf(os.Stderr, "Date: %s\n", resp.Header["Date"])
		fmt.Fprintf(os.Stderr, "Content-Type: %s\n", resp.Header["Content-Type"])
		fmt.Fprintf(os.Stderr, "Content-Length: %d bytes\n", b)
	}
}
