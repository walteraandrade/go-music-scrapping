package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	baseUrl := "https://www.allmusic.com/search/all/afrobeat"

	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	collector.OnHTML("div.album", func(e *colly.HTMLElement) {
		albumName := e.ChildText("div.title a")
		albumURL := e.ChildAttr("div.title a", "href")

		artistName := e.ChildText("div.artist a")
		artistURL := e.ChildAttr("div.artist a", "href")

		year := e.ChildText("div.year")

		fmt.Printf("--- Found Album ---\n")
		fmt.Printf("Album: %s (URL: %s)\n", albumName, albumURL)
		fmt.Printf("Artist: %s (URL: %s)\n", artistName, artistURL)
		fmt.Printf("Year: %s\n", year)
		fmt.Println("-------------------")

		if albumURL != "" {
			fmt.Printf("Visiting individual album page: %s\n", e.Request.AbsoluteURL(albumURL))
			collector.Visit(e.Request.AbsoluteURL(albumURL))
		}
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request URL: %s failed with status code %d and error: %s\n", r.Request.URL.String(), r.StatusCode, err)
	})

	fmt.Printf("Starting scrape from: %s\n", baseUrl)
	collector.Visit(baseUrl)

	collector.Wait()
}
