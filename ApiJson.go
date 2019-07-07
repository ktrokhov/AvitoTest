package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type DB struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
}

var items []DB
//все, что есть в DB
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
// достаем по  id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&DB{})
}
//создавать items
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book DB
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	items = append(items, book)
	json.NewEncoder(w).Encode(book)
}
var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// создаем рандомные имена для каждого id
func srand(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}
func creatingItems() {
	var i int
	for i = 0; i < 2000; i++ {
		items = append(items, DB{ID: srand(2), Name: srand(9)})
	}
}


func main() {
	r := mux.NewRouter()
	creatingItems()
	r.HandleFunc("/api/retrieve/", getBooks).Methods("GET")
	r.HandleFunc("/api/retrieve/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/generate/", createBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", r))
}