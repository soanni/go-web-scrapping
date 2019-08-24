package main

import (
	
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("emojipedia.org"),
	)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		isEmojiPage := false

		metaTags := e.DOM.ParentsUntil("~").Find("meta")
		metaTags.Each(func(_ int, s *goquery.Selection) {
			property, _ := s.Attr("property")
			if strings.EqualFold(property, "og:type") {
				content, _ := s.Attr("content")

				isEmojiPage = strings.EqualFold(content, "article")
			}
		})

		if isEmojiPage {
			fmt.Println("Emoji: ", e.DOM.Find("h1").Text())

			fmt.Println("Decription: ", e.DOM.Find(".description").Find("p").Text())
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