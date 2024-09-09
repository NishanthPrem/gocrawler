package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No website provided")
		return
	}

	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]
	inputMaxConcurrency := os.Args[2]
	intputMaxPages := os.Args[3]

	maxConcurrency, err := strconv.Atoi(inputMaxConcurrency)
	if err != nil {
		fmt.Printf("Invalid maxConcurrency value: %v\n", err)
		return
	}

	maxPages, err := strconv.Atoi(intputMaxPages)
	if err != nil {
		fmt.Printf("Invalid maxPages value: %v\n", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
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
