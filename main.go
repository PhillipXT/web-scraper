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

	const maxConcurrency = 5
	cfg, err := configure(baseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error in configure: %v", err)
		return
	}

	fmt.Println("starting crawl of: ", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(args[1])
	cfg.wg.Wait()

	fmt.Println("Found pages:")
	for page, count := range cfg.pages {
		fmt.Printf("    %d - %s\n", count, page)
	}
}
