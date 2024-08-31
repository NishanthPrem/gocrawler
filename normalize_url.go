package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawUrl string) (string, error) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return "", fmt.Errorf("error parsing the url %w", err)
	}
	normalizedURL := strings.TrimPrefix(url.String(), url.Scheme+"://")
	return normalizedURL, nil
}
