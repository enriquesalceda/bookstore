package bookstore

import "fmt"

type Book struct {
	ID                 string
	Title              string
	Copies             int
	Author             string
	Edition            string
	OnSpecial          bool
	PercentageDiscount int
	Price              int
}

type Catalog map[string]Book

func (c Catalog) GetAllBooks() []Book {
	books := []Book{}

	for _, b := range c {
		books = append(books, b)
	}

	return books
}

func (c *Catalog) AddBook(b Book) {
	(*c)[b.ID] = b
}

func (c Catalog) GetBookDetails(id string) string {
	return c[id].Details()
}

func (b Book) Details() string {
	return fmt.Sprintf("%s, by %s", b.Title, b.Author)
}

func (b Book) NetPrice() int {
	discount := b.Price / b.PercentageDiscount
	return b.Price - discount
}

func (b Book) SalePrice() int {
	discount := b.Price / 2
	return b.Price - discount
}

func (b *Book) SetTitle(t string) {
	b.Title = t
}

func (b *Book) SetPriceCents(p int) {
	b.Price = p
}
