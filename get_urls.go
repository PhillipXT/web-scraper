package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Added a package to make this work:  go get golang.org/x/net/html
func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	var links []string

	doc, err := html.Parse((strings.NewReader(htmlBody)))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, a := range node.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						return nil, fmt.Errorf("couldn't parse href '%v': %v", a.Val, err)
					}
					resolvedURL := baseURL.ResolveReference(href)
					links = append(links, resolvedURL.String())
				}
			}
		}
	}

	return links, nil
}
