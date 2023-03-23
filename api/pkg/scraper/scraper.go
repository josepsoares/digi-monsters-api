package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func ScrapeDigimons() {
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}

func ScrapeDigimonAttributes() {
	c := colly.NewCollector(
		colly.AllowedDomains("wikimon.net"),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("table.wikitable tbody tr", func(e *colly.HTMLElement) {
		// for each to get td
	})

	c.Visit("https://wikimon.net/Attribute")
}

func ScrapeDigimonTypes() {
	c := colly.NewCollector(
		colly.AllowedDomains("wikimon.net"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("table.wikitable tbody tr", func(e *colly.HTMLElement) {
		// for each to get td
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://wikimon.net/Type")
}

func ScrapeDigimonLevels() {
	// https://wikimon.net/Evolution_Stage

	c := colly.NewCollector(
		colly.AllowedDomains("wikimon.net"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://wikimon.net/Evolution_Stage")
}
