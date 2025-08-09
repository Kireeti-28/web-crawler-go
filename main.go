package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseUrl)
	htmlBody, err := getHTML(baseUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(htmlBody)
}
