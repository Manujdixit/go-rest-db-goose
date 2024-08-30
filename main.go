package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// This is a global variable that is used to store the items in the database.
var db *sql.DB

func main() {
	//initialize db connection
	var err error
	db, err = sql.Open("mysql", "root:your_password@tcp(localhost:3306)/itemsdb")
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	defer db.Close()

	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/items/", itemHandler)
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	//get all items
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		//fetch all items from db
		rows, err := db.Query("SELECT * FROM items")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var items []Item
		//This line of code is part of a loop that iterates over the rows returned by a SQL query. The rows. Next() function returns true if there is another row to process, and false otherwise. The loop will continue to execute as long as there are more rows to process.
		for rows.Next() {
			var item Item
			if err := rows.Scan(&item.ID, &item.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}
		json.NewEncoder(w).Encode(items)

	//create new item
	case http.MethodPost:
		//This line of code declares a new variable newItem of type Item, which is a struct defined earlier in the code. The Item struct has two fields: ID and Name. This variable is used to hold a new item that will be created or updated in the code that follows.
		var newItem Item
		//to decode the request body into the newItem variable. It uses the json.NewDecoder function to create a new JSON decoder and then uses the Decode method to decode the request body into the newItem variable. If there is an error during the decoding process, it sends an HTTP error response with a status code of 400 (Bad Request) and the error message. The return statement is used to exit the function if there is an error.
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//insret new item into db
		result, err := db.Exec("INSERT INTO items (name) VALUES (?)", newItem.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//get id of new item
		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve item ID", http.StatusInternalServerError)
			return
		}
		//set id of new item
		newItem.ID = int(id)
		//send response with new item
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newItem)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	//This line of code extracts the ID from the URL path in the itemHandler function. It removes the prefix "/items/" from the URL path and assigns the remaining string to the idStr variable.
	idStr := r.URL.Path[len("/items/"):]
	//converts a string (idStr) to an integer (id) using the strconv.Atoi
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {

	//get item by id
	case http.MethodGet:
		var item Item
		err := db.QueryRow("SELECT * FROM items WHERE id = ?", id).Scan(&item.ID, &item.Name)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)

	//update item by id
	case http.MethodPut:
		var updatedItem Item
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//update the item in the database
		_, err := db.Exec("UPDATE items SET name = ? WHERE id = ?", updatedItem.Name, id)
		if err != nil {
			http.Error(w, "Failed to update item", http.StatusInternalServerError)
			return
		}
		updatedItem.ID = id
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedItem)

	//delete item by id
	case http.MethodDelete:
		_, err := db.Exec("DELETE FROM items WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Failed to delete item", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
