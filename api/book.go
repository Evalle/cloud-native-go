package api

import "encoding/json"

// Book type with Title, Author and ISBN
type Book struct {
	Title  string
	Author string
	ISBN   string
}

// ToJSON to be used for marshalling of Book Type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}
