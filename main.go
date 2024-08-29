package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	link := parseArgs()
	fmt.Println(getHTML(link))

}

func parseArgs() string {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Println("starting crawl of:", args[0])
	return args[0]
}

func getHTML(rawURL string) (string, error) {
	req, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error fetching the data %w", err)
	}

	defer req.Body.Close()

	if req.StatusCode > 399 {
		return "", fmt.Errorf("request failed with status %d: %s", req.StatusCode, http.StatusText(req.StatusCode))
	}

	contentType := req.Header.Get("Content-Type")
	if contentType != "text/html" {
		return "", fmt.Errorf("unexpected content type: %s", contentType)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body %w", err)
	}

	return string(body), nil
}
