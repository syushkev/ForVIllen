package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func getHref(t html.Token) (ok bool, title, href string) {
	for _, a := range t.Attr {
		if a.Key == "title" {
			for _, b := range t.Attr {
				if b.Key == "href" {
					title = a.Val
					href = b.Val
					ok = true
				}
			}
		}
	}

	return
}

func crawl(url, version string) map[string]bool {
	foundUrls := make(map[string]bool)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return foundUrls
	}

	b := resp.Body
	defer b.Close()

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return foundUrls
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			ok, title, url := getHref(t)
			if !ok {
				continue
			}

			if title == version{
				foundUrls[url] = true
			}
		}
	}
	return foundUrls
}

func main() {

	seedUrls := os.Args[1]
	version := os.Args[2]

	m := crawl(seedUrls, version)

	fmt.Println("\nFound", len(m), "unique urls:\n")

	for url, _ := range m {
		fmt.Println(" - " + url)
	}
}

