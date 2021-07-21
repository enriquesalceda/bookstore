package bookstore

import "fmt"

type Book struct {
	ID                 string
	Title              string
	Copies             int
	Author             string
	Edition            string
	OnSpecial          bool
	PercentageDiscount float64
	Price              int
}

type Catalog map[string]Book

func GetAllBooks(catalog Catalog) []Book {
	books := []Book{}

	for _, b := range catalog {
		books = append(books, b)
	}

	return books
}

func GetBookDetails(catalog Catalog, id string) string {
	return catalog[id].Details()
}

func (b Book) Details() string {
	return fmt.Sprintf("%s, by %s", b.Title, b.Author)
}
