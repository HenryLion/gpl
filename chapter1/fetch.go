package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, "http://") {
		} else {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// exercise 1.7
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Copy: %v", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "%s\n", resp.Status)
		/*
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v]n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		*/
	}
	secs := time.Since(start).Seconds()
	fmt.Println("use time: %v", secs)
}
