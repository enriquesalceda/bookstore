package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = map[string]bookstore.Book{
	"1": {
		ID:                 "1",
		Title:              "Hannibal",
		Copies:             1000,
		Author:             "Thomas Harris",
		Edition:            "First edition",
		OnSpecial:          true,
		PercentageDiscount: 0.10,
		Price:              11000,
	},
	"2": {
		ID:                 "2",
		Title:              "Silence of the lambs",
		Copies:             2000,
		Author:             "Thomas Harris",
		Edition:            "Fourth edition",
		OnSpecial:          true,
		PercentageDiscount: 0.05,
		Price:              10000,
	},
}

func TestBook(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:                 "1",
		Title:              "Hannibal",
		Copies:             1000,
		Author:             "Thomas Harris",
		Edition:            "First edition",
		OnSpecial:          true,
		PercentageDiscount: 0.10,
		Price:              10000,
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
			ID:                 "1",
			Title:              "Hannibal",
			Copies:             1000,
			Author:             "Thomas Harris",
			Edition:            "First edition",
			OnSpecial:          true,
			PercentageDiscount: 0.10,
			Price:              11000,
		},
		{
			ID:                 "2",
			Title:              "Silence of the lambs",
			Copies:             2000,
			Author:             "Thomas Harris",
			Edition:            "Fourth edition",
			OnSpecial:          true,
			PercentageDiscount: 0.05,
			Price:              10000,
		},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	want := "Hannibal, by Thomas Harris"
	got := bookstore.GetBookDetails(catalog, "1")

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestBookDetails(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:                 "1",
		Title:              "Hannibal",
		Copies:             1000,
		Author:             "Thomas Harris",
		Edition:            "First edition",
		OnSpecial:          true,
		PercentageDiscount: 0.10,
		Price:              10000,
	}

	want := "Hannibal, by Thomas Harris"
	got := b.Details()

	if want != got {
		t.Errorf("Want: %s, Got: %s", want, got)
	}
}
