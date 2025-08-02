package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error in crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error in crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
	}

	// Don't follow external links
	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	page_url, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error in normalizeURL: %v", err)
		return
	}

	if _, exists := pages[page_url]; exists {
		pages[page_url]++
		fmt.Println("Page already exists, skipping: ", page_url)
		return
	}

	fmt.Println("Added url to list: ", page_url)
	pages[page_url] = 1

	fmt.Println("Reading HTML from ", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error in getHTML: %v\n", err)
		return
	}

	url_list, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("Error in getURLsFromHTML %v", err)
		return
	}

	for _, url := range url_list {
		fmt.Println("Crawling ", url)
		crawlPage(rawBaseURL, url, pages)
		fmt.Printf("URL Count: %v\n", len(pages))
	}
}
