package api

import (
	"testing"
)

func TestBook(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "01234567890"}
	got := string(book.ToJSON())
	want := `{"Title":"Cloud Native Go","Author":"M.-L. Reimer","ISBN":"01234567890"}`

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
