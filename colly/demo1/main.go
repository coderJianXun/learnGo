package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

func main() {
	// 实例化
	c := colly.NewCollector(
		// 限定域名
		colly.AllowedDomains("blog.phpha.com"),
		// 最大深度
		colly.MaxDepth(1),
	)

	for {
		// On every a element which has href attribute call callback
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			// Print link
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			// Visit link found on page
			// Only those links are visited which are in AllowedDomains
			c.Visit(e.Request.AbsoluteURL(link))
		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		// Start scraping on http://blog.phpha.com/
		go func() {
			c.Visit("http://blog.phpha.com/")
		}()
		time.Sleep(1 * time.Second)
	}
}
