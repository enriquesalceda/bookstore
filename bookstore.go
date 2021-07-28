package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID                 string
	Title              string
	Copies             int
	Author             string
	Edition            string
	OnSpecial          bool
	percentageDiscount int
	Price              int
	category           string
}

const (
	Autobiography             = "Autobiography"
	Romance                   = "Romance"
	ScienceFiction            = "ScienceFiction"
	Programming               = "Programming"
	MinimumPercentageDiscount = 0
	MaximumPercentageDiscount = 100
)

var ValidCategories = map[string]bool{
	Autobiography:  true,
	Romance:        true,
	ScienceFiction: false,
	Programming:    true,
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
	discount := b.Price / b.DiscountPercent()
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

func (b *Book) SetDiscountPercent(p int) error {
	if p > MaximumPercentageDiscount || p < MinimumPercentageDiscount {
		return errors.New("discount price should be between 0 and 100")
	}

	b.percentageDiscount = p
	return nil
}

func (b Book) DiscountPercent() int {
	return b.percentageDiscount
}

func TestIsValidCategory(c string) bool {
	return ValidCategories[c]
}

func (b Book) GetCategory() string {
	return b.category
}

func (b *Book) SetCategory(c string) error {
	if ValidCategories[c] {
		b.category = c
		return nil
	}

	return errors.New("this is not a valid category")
}
