package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"astuart.co/goq"
	"github.com/gocolly/colly/v2"
)

type Company struct {
	Name               string `goquery:"div.startup-company > div.name"                                             json:"name"`
	Logo               string `goquery:"div.startup-company > div.thumb > img,[src]"                                json:"logo"`
	Representative     string `goquery:"div.startup-company-detail > div:nth-child(1) > dl:nth-child(1) > dd"       json:"representative"`
	Location           string `goquery:"div.startup-company-detail > div:nth-child(1) > dl:nth-child(2) > dd"       json:"location"`
	EstablishedAt      string `goquery:"div.startup-company-detail > div:nth-child(2) > dl:nth-child(1) > dd"       json:"establishedAt"`
	Link               string `goquery:"div.startup-company-detail > div:nth-child(2) > dl:nth-child(2) > dd"       json:"link"`
	Mail               string `goquery:"div.startup-company-detail > div:nth-child(3) > dl:nth-child(1) > dd"       json:"mail"`
	Telephone          string `goquery:"div.startup-company-detail > div:nth-child(3) > dl:nth-child(2) > dd"       json:"telephone"`
	Domain             string `goquery:"div.startup-company-detail > div:nth-child(4) > dl:nth-child(1) > dd"       json:"domain"`
	MainProduct        string `goquery:"div.startup-company-detail > div:nth-child(4) > dl:nth-child(2) > dd"       json:"mainProduct"`
	CLevel             string `goquery:"div.startup-company-detail > div:nth-child(5) > dl:nth-child(1) > dd"       json:"cLevel"`
	Employees          string `goquery:"div.startup-company-detail > div:nth-child(5) > dl:nth-child(2) > dd"       json:"employees"`
	Investment         string `goquery:"div.startup-company-detail > div.item-row.type-line > dl:nth-child(1) > dd" json:"investment"`
	Series             string `goquery:"div.startup-company-detail > div.item-row.type-line > dl:nth-child(2) > dd" json:"series"`
	InvestmentOverview string `goquery:"div.startup-company-detail > div:nth-child(7) > dl > dd"                    json:"investmentOverview"`
	Investor           string `goquery:"div.startup-company-detail > div:nth-child(8) > dl > dd"                    json:"investor"`
}

const nextSelector = `#container > div > div.db-search-result > div.db-search-list > div > a.btn-page-next`

const detailSelector = `#container > div > div.db-search-result > div.db-search-list > ul > li > div.txt-cont > div.startup-name > a`

const detailBoxSelector = `#container > div.box.startup-db-view`

var limit = &colly.LimitRule{
	DomainGlob:  "*",
	Parallelism: 1,
	Delay:       time.Millisecond * 10,
}

func retry(resp *colly.Response, err error) {
	log.Println(err)

	resp.Request.Retry()
}

const baseUrl = "https://www.hankyung.com"

const entryPoint = baseUrl + "/geeks/startupdb"

func main() {
	filename := flag.Arg(0)
	if filename == "" {
		filename = "companies.ndjson"
	}

	f, err := os.OpenFile(
		filename,
		os.O_APPEND|
			os.O_WRONLY|
			os.O_CREATE,
		0600,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	// enc.SetIndent("", "\t")

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
		companyDetailCollector.Visit(baseUrl + e.Attr("href"))
	})

	companyDetailCollector.OnHTML(detailBoxSelector, func(e *colly.HTMLElement) {
		var company Company
		if err := goq.UnmarshalSelection(e.DOM, &company); err != nil {
			panic(err)
		}

		enc.Encode(company)
	})

	paginationCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Page", r.URL)
	})

	companyDetailCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Company", r.URL)
	})

	paginationCollector.OnError(retry)
	companyDetailCollector.OnError(retry)

	paginationCollector.Visit(entryPoint)

	paginationCollector.Wait()
}
