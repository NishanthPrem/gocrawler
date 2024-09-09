package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	// To control the number of goroutines
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	//Parse the url

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	//Skip all other websites

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	// Normalize the url

	normalizeURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
	}

	// If the url is already visited, dont crawl again
	isFirst := cfg.addPageVisit(normalizeURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	// Crawling the page and getting the URL
	htmlBody, err := getHTML(rawCurrentURL)

	if err != nil {
		fmt.Printf("Error - getHTML %v", err)
	}

	// Getting all the URLs in the particular page and start crawling
	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting url from html %v", err)
	}

	//Run the for loop concurrently
	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
