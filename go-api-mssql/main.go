package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// Define a struct to represent your entity
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *sql.DB

func main() {
	// Connect to the MSSQL database
	var err error
	db, err = sql.Open("sqlserver", "server=localhost;user id=;password=;database=godb;trusted_connection=yes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize HTTP router
	r := mux.NewRouter()

	// Define API endpoints
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Printf("Server started at port 8080")
}

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the item into the database
	_, err := db.Exec("INSERT INTO item (name, price) VALUES (?, ?)", newItem.Name, newItem.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	// Query the database for all items
	rows, err := db.Query("SELECT id, name, price FROM item")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the retrieved items
	var items []Item

	// Iterate through the rows and scan each item into the slice
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	// Check for any errors during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the items slice to JSON
	res, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Get an item by ID
func getItem(w http.ResponseWriter, r *http.Request) {
	// Extract item ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database for the item with the given ID
	var item Item
	err := db.QueryRow("SELECT id, name, price FROM item WHERE id = ?", id).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Serialize item to JSON and send it in the response
	json.NewEncoder(w).Encode(item)
}

// Update an item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	// Extract item ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Parse request body
	var updatedItem Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the item in the database
	_, err := db.Exec("UPDATE item SET name = ?, price = ? WHERE id = ?", updatedItem.Name, updatedItem.Price, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete an item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Extract item ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the item from the database
	_, err := db.Exec("DELETE FROM item WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
