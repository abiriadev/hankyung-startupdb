package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

const nextSelector = `#container > div > div.db-search-result > div.db-search-list > div > a.btn-page-next`

const detailSelector = `#container > div > div.db-search-result > div.db-search-list > ul > li > div.txt-cont > div.startup-name > a`

var limit = &colly.LimitRule{
	DomainGlob:  "*",
	Parallelism: 1,
	Delay:       time.Second,
}

func retry(resp *colly.Response, err error) {
	resp.Request.Retry()
}

func main() {
	paginationCollector := colly.NewCollector(
		colly.Async(),
	)
	companyDetailCollector := colly.NewCollector()

	paginationCollector.Limit(limit)
	companyDetailCollector.Limit(limit)

	paginationCollector.OnHTML(nextSelector, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	paginationCollector.OnHTML(detailSelector, func(e *colly.HTMLElement) {
		companyDetailCollector.Visit(e.Attr("href"))
	})

	paginationCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Page", r.URL)
	})

	companyDetailCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Company", r.URL)
	})

	paginationCollector.OnError(retry)
	companyDetailCollector.OnError(retry)

	paginationCollector.Visit("https://www.hankyung.com/geeks/startupdb")

	paginationCollector.Wait()
}
