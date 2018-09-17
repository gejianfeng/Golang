package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		_url := url

		if !strings.HasPrefix(_url, "http://") {
			_url = "http://" + _url
		}

		resp, err := http.Get(_url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", _url, err)
			os.Exit(1)
		}
	}
}