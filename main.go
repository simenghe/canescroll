package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

const sampleURL = "https://www.reddit.com/r/pics/"
const testURL = "https://www.reddit.com/r/pics/comments/pe6iwq/sharing_this_photo_cause_my_friend_doesnt_have/"

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

func preeImage() {
	c := colly.NewCollector()
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("src"))
	})
	c.Visit(testURL)
}

func main() {
	viableURLs := []string{}
	viableImages := []string{}
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, sampleURL+"comments") {
			e.Request.Visit(e.Attr("href"))
			viableURLs = append(viableURLs, e.Attr("href"))
		}
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println(r.URL)
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println(r.ID)
	// })

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		// fmt.Println("Image Source : ", src)
		if strings.HasPrefix(src, "https://preview.redd.it/award_images") {
			return
		}
		if strings.HasPrefix(src, "https://preview.redd.it") {
			viableImages = append(viableImages, src)
		}
	})

	c.Visit(sampleURL)
	fmt.Println(uniqueImageURLs(viableImages))
}

func uniqueImageURLs(urls []string) []string {
	newUrls := make([]string, len(urls))
	visited := make(map[string]bool)

	for _, uri := range urls {
		parsedURL, err := url.Parse(uri)
		if err != nil {
			return newUrls
		}

		if visited[parsedURL.Path] {
			continue
		}
		newUrls = append(newUrls, parsedURL.String())
		visited[parsedURL.Path] = true
		// fmt.Println(parsedURL.Path)
		// fmt.Println(parsedURL.Host)
		// fmt.Println(parsedURL.Query())
	}
	return newUrls
}
