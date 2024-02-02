package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Book struct {
	Title  string
	Author string
	Votes  int32
}

// This is simple page scraper which will return TOP 100 books from lubimyczytac.pl
func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".row.relative.w-100", func(e *colly.HTMLElement) {
		e.ForEach(".authorAllBooks__singleImgWrap", func(_ int, el *colly.HTMLElement) {
			imageURL := el.ChildAttr("img", "src")
			fmt.Println("Image URL:", imageURL)
		})

		e.ForEach(".authorAllBooks__singleTextAuthor", func(_ int, el *colly.HTMLElement) {
			author_url := el.ChildAttr("a", "href")
			fmt.Println("Autor URL:", author_url)
		})

		// For each nested <div> with class "authorAllBooks__singleText", extract the book title
		e.ForEach(".authorAllBooks__singleText", func(_ int, el *colly.HTMLElement) {
			bookTitle := el.ChildText(".authorAllBooks__singleTextTitle")
			fmt.Println("Book Title:", bookTitle)
		})

		// For each nested <div> with class "listLibrary__info", extract additional information
		e.ForEach(".listLibrary__info", func(_ int, el *colly.HTMLElement) {
			rating := el.ChildText(".listLibrary__ratingStarsNumber")
			ratingsCount := el.ChildText(".listLibrary__ratingAll")
			fmt.Println("Rating:", rating)
			fmt.Println("Ratings Count:", ratingsCount)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err)
	})

	c.Visit("https://lubimyczytac.pl/top100")
}
