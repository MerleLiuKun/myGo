// code for exercise 1.7
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stderr, resp.Body)  // return content_length and err
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Copying %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

// go run fetch_copy.go http://www.baidu.com