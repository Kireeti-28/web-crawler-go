package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, baseRawURL string) ([]string, error) {

	baseURL, err := url.Parse(baseRawURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	htmlNodes, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string

	for n := range htmlNodes.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					// convert relative url to absolute
					hrefURL, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", attr.Val, err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(hrefURL).String()
					urls = append(urls, resolvedURL)
				}
			}
		}
	}

	return urls, nil
}
