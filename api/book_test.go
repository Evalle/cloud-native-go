package api

import (
	"testing"
)

// Positive tests
func TestBook(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "01234567890"}
	got := string(book.ToJSON())
	want := `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"01234567890"}`

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func TestBookFromJson(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"01234567890"}`)
	want := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "01234567890"}
	got := FromJSON(json)

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
