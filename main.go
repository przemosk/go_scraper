package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type Book struct {
	Title     string
	Author    string
	AuthorURL string
	ImageURL  string
	Votes     string
}

func ExtractBookAttributes(e *colly.HTMLElement) map[string]string {
	image_url := e.ChildAttr("img", "src")
	author_name := e.ChildText(".authorAllBooks__singleTextAuthor")
	book_title := e.ChildText(".authorAllBooks__singleTextTitle")
	author_url := e.ChildAttr(".authorAllBooks__singleTextAuthor a", "href")
	rating := e.ChildText(".listLibrary__ratingAll") // split and get only number

	book := make(map[string]string)
	book["image_url"] = image_url
	book["author_name"] = author_name
	book["book_title"] = book_title
	book["author_url"] = author_url
	book["rating"] = rating

	return book
}

func TotalPageNumber() int {
	return 7 // so far, still static
}

func CreateJSONFile(books []Book) {
	jsonData, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	file, err := os.Create("books.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("\nFILE CREATED: ", file.Name())
}

func main() {
	var books []Book

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		// log.Print("VISITING ", r.URL)
	})

	c.OnHTML(".container", func(e *colly.HTMLElement) {
		e.ForEach(".authorAllBooks__single", func(_ int, e *colly.HTMLElement) {
			book_attribute := ExtractBookAttributes(e)

			book := Book{
				Title:     book_attribute["book_title"],
				Author:    book_attribute["author_name"],
				AuthorURL: book_attribute["author_url"],
				ImageURL:  book_attribute["image_url"],
				Votes:     book_attribute["rating"]}

			books = append(books, book)
			log.Print(book)
		})

	})

	c.OnError(func(r *colly.Response, err error) {
		log.Fatal(err)
	})

	for i := 1; i <= TotalPageNumber(); i++ {
		url := fmt.Sprintf("https://lubimyczytac.pl/top100?page=%d", i)
		c.Visit(url)
	}

	CreateJSONFile(books)
}
