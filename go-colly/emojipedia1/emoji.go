package main

import (
	
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("emojipedia.org"),
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		if strings.EqualFold(e.ChildAttr(`meta[property="og:type"]`, "content"), "article") {
			fmt.Println("Emoji: ", e.ChildText("article h1"))
			fmt.Println("Decription: ", e.ChildText("article .description p"))
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Limit(&colly.LimitRule{
		DomainGlob: "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://emojipedia.org")
}