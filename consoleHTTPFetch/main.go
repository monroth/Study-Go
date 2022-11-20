package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
)

func downloadHTML(site string) {
	start := time.Now()

	res, err := client.Get(site)
	if err != nil {
		fmt.Printf("%s Got http.Get error \"%s\" after waiting for %s\n", site, err, time.Since(start))
		return
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("%s Failed respone with status code %d after waiting for %s\n", site, res.StatusCode, time.Since(start))
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s Got io.ReadAll error \"%s\" after waiting for %s\n", site, err, time.Since(start))
		return
	}
	fmt.Printf("%s Got response of size %d after waiting for %s\n", site, len(body), time.Since(start))

}

func main() {
	var wg sync.WaitGroup
	args := os.Args[1:]
	for _, arg := range args {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			downloadHTML(url)
		}(arg)
	}
	wg.Wait()
}
