package main

import (
	"net/url"
	"strings"
)

func normalizeURL(input string) (string, error) {
	url, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	urlPath := url.Path
	if urlPath == "" || urlPath == "/" {
		return url.Hostname(), nil
	}

	urlPath, _ = strings.CutSuffix(urlPath, "/")

	return strings.ToLower(url.Hostname() + urlPath), nil
}
