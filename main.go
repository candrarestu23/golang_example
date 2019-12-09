package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int    `jsonL:"id"`
	Title  string `jsonL:"title"`
	Author string `jsonL:"author"`
	Year   string `jsonL:"year"`
}

var books []Book
var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()
}

func main() {
	// connStr := "user=" + os.Getenv("username") +
	// 	" dbname=" + os.Getenv("dbName") +
	// 	" password=" + os.Getenv("password") +
	// 	" host=" + os.Getenv("URL") +
	// 	" sslmode=disable"

	connStr := "user=candra" +
		" dbname=testdb" +
		" password=candra1995" +
		" host=167.71.200.137" +
		" sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books = []Book{}

	rows, err := db.Query("select * from test_table")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	print(3)
}

func addBook(w http.ResponseWriter, r *http.Request) {
}

func updateBook(w http.ResponseWriter, r *http.Request) {
}

func removeBook(w http.ResponseWriter, r *http.Request) {
}
