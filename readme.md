# Web Crawler (Go)

A simple concurrent web crawler written in Go. It crawls a website, collects internal links, and prints a report of the number of times each internal link was found.

## Features

- Concurrent crawling with configurable concurrency
- Limits the number of pages crawled
- Normalizes URLs for accurate internal link counting
- Parses HTML and extracts links robustly
- Prints a summary report of internal links

## Usage

```sh
go run main.go <baseURL> <maxConcurrency> <maxPages>
```

- `<baseURL>`: The starting URL to crawl (e.g., `https://example.com`)
- `<maxConcurrency>`: Maximum number of concurrent requests (e.g., `5`)
- `<maxPages>`: Maximum number of pages to crawl (e.g., `100`)

Example:

```sh
go run main.go https://blog.boot.dev 5 100
```

## Project Structure

- `main.go` - Entry point, argument parsing, and crawl orchestration
- `config.go` - Configuration and state management
- `crawl_page.go` - Page crawling logic
- `get_html.go` - Fetches HTML content from URLs
- `get_urls_from_html.go` - Extracts URLs from HTML
- `normalize_url.go` - Normalizes URLs for internal link tracking
- `print_report.go` - Prints the crawl report
- `*_test.go` - Unit tests for core functions

## Running Tests

```sh
go test ./...
```

## Requirements

- Go 1.24.3 or newer