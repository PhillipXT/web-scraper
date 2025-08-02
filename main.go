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
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[1]

	fmt.Println("starting crawl of: ", baseURL)

	pages := map[string]int{}
	crawlPage(baseURL, baseURL, pages)

	fmt.Println("Found pages:")
	for page, count := range pages {
		fmt.Printf("    %d - %s\n", count, page)
	}

}
