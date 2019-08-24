package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		fmt.Println("UserAgent", r.Headers.Get("User-Agent"))
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var pName, stars, price string

			pName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")
			
			stars = e.ChildText("span.a-icon-alt")
			// FormatStars(&stars)

			price = e.ChildText("span.a-price > span.a-offscreen")
			// FormatPrice(&price)

			if pName == "" {
				return
			}

			fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", pName, stars, price)
		})
	})

	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})

	for i := 1; i <= 20; i++ {
		c.Visit(fmt.Sprintf("https://www.amazon.com/s?k=nintendo+switch&page=%d", i))
	}
	
	c.Wait()	
}

func FormatPrice(price *string) {
	r := regexp.MustCompile(`\$(\d+(\.\d+)?).*$`)

	newPrices := r.FindStringSubmatch(*price)

	if len(newPrices) > 1 {
		*price = newPrices[1]
	} else {
		*price = "Unknown"
	}
}

func FormatStars(stars *string) {
	if len(*stars) >=3 {
		*stars = (*stars)[0:3]
	} else {
		*stars = "Unknown"
	}
}