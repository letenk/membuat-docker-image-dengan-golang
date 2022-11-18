package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const serverPort = ":3000"

// Struct Book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Sample data
var books = []Book{
	{ID: 1, Title: "REWORK", Author: "Jason Fried & David Heinemeier Hansson"},
	{ID: 2, Title: "Atomic Habbits", Author: "James Clear"},
}

func main() {
	// Route Home
	http.HandleFunc("/", logMiddleware(getHome))
	// Route Get Books
	http.HandleFunc("/books", logMiddleware(getBooks))
	// Route Get Book By id
	http.HandleFunc("/book", logMiddleware(getBook))
	// Route Post Book
	http.HandleFunc("/post-book", logMiddleware(postBook))

	// Print server is starts
	fmt.Printf("Server starting at http://localhost:%s\n", serverPort)

	// Create new server
	http.ListenAndServe(serverPort, nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	// message response
	message := []byte("Home page")
	// write response
	w.Write(message)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, _ := strconv.Atoi(r.FormValue("id"))

		for _, data := range books {
			if data.ID == id {
				json.NewEncoder(w).Encode(data)
				return
			}
		}

		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func postBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newBook Book
		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			log.Panic(err)
		}

		books = append(books, newBook)

		w.Write([]byte("Book has been added"))
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Address: %s, Url: %s, Method: %s\n", r.RemoteAddr, r.URL, r.Method)
		next.ServeHTTP(w, r)
	}
}
