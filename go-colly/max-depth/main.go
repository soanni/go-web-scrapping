package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		e.Request.Visit(link)
	})

	c.Visit("https://en.wikipedia.org/")
}