package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.numPages() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error in crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
	}

	// Don't follow external links
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	page_url, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error in normalizeURL: %v", err)
		return
	}

	if !cfg.addPageVisit(page_url) {
		return
	}

	//fmt.Println("Reading HTML from ", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error in getHTML: %v\n", err)
		return
	}

	url_list, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error in getURLsFromHTML %v", err)
		return
	}

	for _, url := range url_list {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
