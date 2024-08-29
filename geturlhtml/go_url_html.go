package gourlhtml

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing url %w", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("error reading html bodu %w", err)
	}

	var urls []string
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					u, err := url.Parse(attr.Val)
					if err == nil {
						resolved := base.ResolveReference(u)
						urls = append(urls, resolved.String())
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	return urls, nil

}
