package order_test

import (
	"bookstore/order"
	"testing"
)

func TestOrderNew(t *testing.T) {
	_, err := order.New("111111111111111")

	if err != nil {
		t.Error("Not working")
	}
}

func TestGetId(t *testing.T) {
	want := "111111111111111"
	order, _ := order.New(want)

	got := order.GetId()

	if got != want {
		t.Errorf("Want: %s, Got: %s", want, got)
	}
}

func TestSetId(t *testing.T) {
	want := "2"
	order, _ := order.New("1")
	order.SetId(want)

	got := order.GetId()

	if got != want {
		t.Errorf("Want: %s, Got: %s", want, got)
	}
}
