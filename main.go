package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

const next = `#container > div > div.db-search-result > div.db-search-list > div > a.btn-page-next`

func main() {
	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		Delay:       time.Second,
	})

	c.OnHTML(next, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(resp *colly.Response, err error) {
		resp.Request.Retry()
	})

	c.Visit("https://www.hankyung.com/geeks/startupdb")
}
