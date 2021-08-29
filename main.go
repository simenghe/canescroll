package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

const sampleURL = "https://www.reddit.com/r/pics/"

func getCommentURLs() []string {
	viableURLs := []string{}
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		if strings.HasPrefix(link, sampleURL+"comments") {
			viableURLs = append(viableURLs, link)
		}
	})
	c.Visit(sampleURL)
	return viableURLs
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, sampleURL+"comments") {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
		err := c.Visit(r.URL.String())
		fmt.Println(err)
	})

	c.Visit(sampleURL)
}
