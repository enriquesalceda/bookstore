package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var catalog = bookstore.Catalog{
	"1": {
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     11000,
	},
	"2": {
		ID:        "2",
		Title:     "Silence of the lambs",
		Copies:    2000,
		Author:    "Thomas Harris",
		Edition:   "Fourth edition",
		OnSpecial: true,
		Price:     10000,
	},
}

func TestBook(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     10000,
	}

	want := "Hannibal"
	got := b.Title

	if want != got {
		t.Error("Not expected title")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	want := []bookstore.Book{
		{
			ID:        "1",
			Title:     "Hannibal",
			Copies:    1000,
			Author:    "Thomas Harris",
			Edition:   "First edition",
			OnSpecial: true,
			Price:     11000,
		},
		{
			ID:        "2",
			Title:     "Silence of the lambs",
			Copies:    2000,
			Author:    "Thomas Harris",
			Edition:   "Fourth edition",
			OnSpecial: true,
			Price:     10000,
		},
	}

	got := catalog.GetAllBooks()

	fmt.Println(want)
	fmt.Println(got)

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	want := "Hannibal, by Thomas Harris"
	got := catalog.GetBookDetails("1")

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestBookDetails(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     10000,
	}

	want := "Hannibal, by Thomas Harris"
	got := b.Details()

	if want != got {
		t.Errorf("Want: %s, Got: %s", want, got)
	}
}

func TestNetPrice(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     10000,
	}

	b.SetDiscountPercent(10)
	want := 9000
	got := b.NetPrice()

	if want != got {
		t.Errorf("Want net price: %d, Got net price: %d", want, got)
	}
}

func TestSalePrice(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     10000,
	}

	want := 5000
	got := b.SalePrice()

	if want != got {
		t.Errorf("Want sale price: %d, Got sale price: %d", want, got)
	}
}

func TestSetTitle(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     10000,
	}

	want := "Hannibal Lecter"
	b.SetTitle("Hannibal Lecter")
	got := b.Title

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     10000,
	}

	b.SetPriceCents(8000)
	want := 8000
	got := b.Price

	if got != want {
		t.Errorf("Want price: %d, Got price: %d", want, got)
	}
}

func TestAddBook(t *testing.T) {
	book := bookstore.Book{
		ID:        "3",
		Title:     "Watchmen",
		Copies:    3000,
		Author:    "Alan Moore",
		Edition:   "6th edition",
		OnSpecial: false,
		Price:     20000,
	}

	catalog := bookstore.Catalog{
		"1": {
			ID:        "1",
			Title:     "Hannibal",
			Copies:    1000,
			Author:    "Thomas Harris",
			Edition:   "First edition",
			OnSpecial: true,
			Price:     11000,
		},
		"2": {
			ID:        "2",
			Title:     "Silence of the lambs",
			Copies:    2000,
			Author:    "Thomas Harris",
			Edition:   "Fourth edition",
			OnSpecial: true,
			Price:     10000,
		},
	}

	catalog.AddBook(book)

	want := bookstore.Catalog{
		"1": {
			ID:        "1",
			Title:     "Hannibal",
			Copies:    1000,
			Author:    "Thomas Harris",
			Edition:   "First edition",
			OnSpecial: true,
			Price:     11000,
		},
		"2": {
			ID:        "2",
			Title:     "Silence of the lambs",
			Copies:    2000,
			Author:    "Thomas Harris",
			Edition:   "Fourth edition",
			OnSpecial: true,
			Price:     10000,
		},
		"3": {
			ID:        "3",
			Title:     "Watchmen",
			Copies:    3000,
			Author:    "Alan Moore",
			Edition:   "6th edition",
			OnSpecial: false,
			Price:     20000,
		},
	}

	got := catalog

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetDiscountPercent(t *testing.T) {
	book := bookstore.Book{
		ID:        "3",
		Title:     "Watchmen",
		Copies:    3000,
		Author:    "Alan Moore",
		Edition:   "6th edition",
		OnSpecial: false,
		Price:     20000,
	}

	errorOne := book.SetDiscountPercent(200)

	if errorOne == nil {
		t.Errorf("Expected error if discount percent greather than 100, got: %s", errorOne)
	}

	errorTwo := book.SetDiscountPercent(-100)

	if errorTwo == nil {
		t.Errorf("Expected error if discount percent greather than 100, got: %s", errorTwo)
	}

	book.SetDiscountPercent(20)

	got := book.DiscountPercent()
	want := 20

	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}
}

func TestValidCategories(t *testing.T) {
	tsc := []struct {
		category string
		want     bool
	}{
		{
			category: bookstore.Autobiography,
			want:     true,
		},
		{
			category: bookstore.Romance,
			want:     true,
		},
		{
			category: bookstore.ScienceFiction,
			want:     false,
		},
		{
			category: bookstore.Programming,
			want:     true,
		},
	}

	for _, ts := range tsc {
		got := bookstore.ValidCategories[ts.category]
		want := ts.want

		if want != got {
			t.Errorf("Want: %v, Got: %v", want, got)
		}
	}
}

func TestCategoryAccessors(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:        "1",
		Title:     "Hannibal",
		Copies:    1000,
		Author:    "Thomas Harris",
		Edition:   "First edition",
		OnSpecial: true,
		Price:     10000,
	}

	initialCategory := b.GetCategory()

	if initialCategory != "" {
		t.Errorf("Initial category must be an empty string")
	}

	err := b.SetCategory(bookstore.Programming)

	if err != nil {
		t.Error("Setting bad category should raise an error")
	}

	want := bookstore.Autobiography
	b.SetCategory(want)
	got := b.GetCategory()

	if want != got {
		t.Errorf("Expected: %s category. Received: %s", want, got)
	}
}
