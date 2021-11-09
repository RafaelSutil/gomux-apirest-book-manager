package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct do Livro (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Estrutura do Author
type Author struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

// Init Books variable as a slice book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get params
	params := mux.Vars(r)
	for _, v := range books {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, v := range books {
		if v.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, v := range books {
		if v.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	fmt.Println("Hello World")
	// Inicializa roteador
	r := mux.NewRouter()

	MockData()

	// Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func MockData() {
	//Mock Data - @todo - implementar DB
	books = append(books,
		Book{ID: "1", Isbn: "123321", Title: "Pet Sematary ",
			Author: &Author{
				Firstname: "Stephen",
				Lastname:  "King",
			}})
	books = append(books,
		Book{ID: "2", Isbn: "321123", Title: "The Green Mile",
			Author: &Author{
				Firstname: "Stephen",
				Lastname:  "King",
			}})
	books = append(books,
		Book{ID: "3", Isbn: "789987", Title: "Harry Potter and the Sorcerer's Stone",
			Author: &Author{
				Firstname: "Joanne",
				Lastname:  "Rowling",
			}})
	books = append(books,
		Book{ID: "4", Isbn: "987789", Title: "A Maior Boca do Mundo",
			Author: &Author{
				Firstname: "Lúcia",
				Lastname:  "Góes",
			}})
	books = append(books,
		Book{ID: "5", Isbn: "456654", Title: "Clean Code",
			Author: &Author{
				Firstname: "Robert",
				Lastname:  "Martin",
			}})
}
