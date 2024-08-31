package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No website provided")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]

	fmt.Printf("Starting to crawl %s \n", rawBaseURL)

	pages := make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizeURL, count := range pages {
		fmt.Printf("%d - $s\n", count, normalizeURL)
	}
}
