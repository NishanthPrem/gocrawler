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

	const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("error with configure %v", err)
	}

	fmt.Printf("Starting crawl of %s....\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}

}
