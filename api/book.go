package api

import (
	"encoding/json"
	"net/http"
)

// Book type with Title, Author and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// Books is a slice of all known books
var Books = []Book{
	Book{Title: "Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	Book{Title: "Cloud Native Go", Author: "M.Reimer", ISBN: "1234567890"},
}

// ToJSON to be used for marshalling of Book Type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return ToJSON
}

// FromJSON to be used for unmarshalling of Book Type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}

	return book
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
