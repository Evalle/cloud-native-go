package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/evalle/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/books", api.BooksHandleFunc)

	http.ListenAndServe(":"+port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("Hello, Cloud Native Go")
}
