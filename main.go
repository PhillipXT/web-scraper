package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	maxPages := 10
	maxConcurrency := 5

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[1]

	if len(args) >= 3 {
		val, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Max concurrency must be an integer value.")
		}
		maxConcurrency = val
	}

	if len(args) == 4 {
		val, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("Max pages must be an integer value.")
		}
		maxPages = val
	}

	cfg, err := configure(baseURL, maxPages, maxConcurrency)
	if err != nil {
		fmt.Printf("Error in configure: %v", err)
		return
	}

	fmt.Println("starting crawl of: ", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(args[1])
	cfg.wg.Wait()

	cfg.printReport()
}
