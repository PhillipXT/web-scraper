package main

import (
	"fmt"
	"net/url"
	"sort"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

type page struct {
	pageURL   string
	pageCount int
}

func (cfg *config) addPageVisit(normalizedURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) numPages() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func (cfg *config) printReport() {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", cfg.baseURL.String())
	fmt.Println("=============================")

	pages := []page{}
	for pageURL, pageCount := range cfg.pages {
		pages = append(pages, page{pageURL: pageURL, pageCount: pageCount})
	}

	sort.Slice(pages, func(i, j int) bool {
		if pages[i].pageCount == pages[j].pageCount {
			return pages[i].pageURL < pages[j].pageURL
		}
		return pages[i].pageCount > pages[j].pageCount
	})

	for _, p := range pages {
		fmt.Printf("Found %d internal links to %s\n", p.pageCount, p.pageURL)
	}
}

func configure(rawBaseURL string, maxPages, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}
