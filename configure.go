package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	//Blocking the thread
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	// Updating the map
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func configure(rawBaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	//Parse the url and return the config struct
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)

	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil

}
