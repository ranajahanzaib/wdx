/**
 * Copyright (c) 2021-present, Rana Jahanzaib
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package wdx

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Log progress / details on the go
func Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Scrap links from a given URL using class name, and return the links
func ScrapLinks(url string, className string) []string {
	var links []string
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML(className, func(e *colly.HTMLElement) {
		links = e.ChildAttrs("a", "href")
	})
	c.Visit(url)
	return links
}
