package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type Book struct {
	Title     string
	Author    string
	AuthorURL string
	ImageURL  string
	Votes     string
}

func main() {
	var books []Book

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Print("VISITING ", r.URL)
	})

	c.OnHTML(".container", func(e *colly.HTMLElement) {
		e.ForEach(".authorAllBooks__single", func(_ int, e *colly.HTMLElement) {
			image_url := e.ChildAttr("img", "src")
			author_name := e.ChildText(".authorAllBooks__singleTextAuthor")
			book_title := e.ChildText(".authorAllBooks__singleTextTitle")
			author_url := e.ChildAttr(".authorAllBooks__singleTextAuthor a", "href")
			rating := e.ChildText(".listLibrary__ratingAll")

			book := Book{Title: book_title, Author: author_name, AuthorURL: author_url, ImageURL: image_url, Votes: rating}
			books = append(books, book)
			log.Print(book)
		})

	})

	c.OnError(func(r *colly.Response, err error) {
		log.Fatal(err)
	})

	var total_page int = 7
	for i := 1; i <= total_page; i++ {
		url := fmt.Sprintf("https://lubimyczytac.pl/top100?page=%d", i)
		c.Visit(url)
	}
}
