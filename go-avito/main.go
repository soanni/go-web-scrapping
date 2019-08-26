package main

import (
	"fmt"
	"time"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	_ "github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(
		// colly.AllowedDomains("avito.ru"),
		// colly.Debugger(&debug.LogDebugger{}),
	)

	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		fmt.Println("UserAgent", r.Headers.Get("User-Agent"))
	})

	c.OnHTML("div.catalog-list.js-catalog-list.clearfix", func(e *colly.HTMLElement) {
		e.ForEach("div.item_table-header", func(_ int, e *colly.HTMLElement) {
			var link, price string
			link = e.ChildAttr("a.item-description-title-link", "href")
			price = e.ChildAttr("span.price.price_highlight", "content")
			fmt.Printf("Link: %s \nPrice: %s \n", link, price)

			c.Visit(e.Request.AbsoluteURL(link))
		})
	})

	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})

	c.Visit("https://www.avito.ru/samara/kvartiry")

	// for i := 1; i <= 20; i++ {
	// 	c.Visit(fmt.Sprintf("https://www.amazon.com/s?k=nintendo+switch&page=%d", i))
	// }
	
	c.Wait()	
}