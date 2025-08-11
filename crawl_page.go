package main

import (
	"fmt"
	"net/url"
)


func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base url: %v\n", err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current url: %v\n", err)
		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizeURL: %v", err)
		return 
	}

	if _, ok := pages[normalizedURL]; ok {
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL]++

	
	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting HTML: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting URLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		crawlPage(rawBaseURL, nextURL, pages)
	}
}