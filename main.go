package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// This is simple page scraper which will return TOP 100 books from lubimyczytac.pl
func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".authorAllBooks__single", func(e *colly.HTMLElement) {
		fmt.Println("in OnHTML", e.Name)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err)
	})

	c.Visit("https://lubimyczytac.pl/top100")

	c.Wait()
}
