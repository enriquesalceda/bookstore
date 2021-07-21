package main

import (
	"fmt"
)

type Book struct {
	Title              string
	Copies             int
	Author             string
	Edition            string
	OnSpecial          bool
	PercentageDiscount float64
	Price              int
}

func (h Book) NormalPrice() float64 {
	return float64(h.Price / 100)
}

func (h Book) CurrentPrice() float64 {
	bookCurrentPrice := h.NormalPrice()

	if h.OnSpecial {
		bookCurrentPrice = bookCurrentPrice - (bookCurrentPrice * h.PercentageDiscount)
	}

	return bookCurrentPrice
}

func formatPrice(p float64) string {
	return fmt.Sprintf("$%.2f", p)
}

func main() {
	book := Book{
		Title:              "Hannibal",
		Copies:             1000,
		Author:             "Thomas Harries",
		Edition:            "First edition",
		OnSpecial:          true,
		PercentageDiscount: 0.10,
		Price:              10000,
	}

	fmt.Println(formatPrice(book.CurrentPrice()))
}
