package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `jsonL:"id"`
	Title  string `jsonL:"title"`
	Author string `jsonL:"author"`
	Year   string `jsonL:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "GoRoutines", Author: "Mr. Goroutines", Year: "2011"},
		Book{ID: 3, Title: "Golang Routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang Concurrency", Author: "Mr. Concurency", Year: "2013"},
		Book{ID: 5, Title: "Golang Good Parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("get Book is Called")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("add Book is Called")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("update Book is Called")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("remove Book is Called")
}