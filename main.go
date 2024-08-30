// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manujdixit/go-rest-db-goose/controllers"
	"github.com/manujdixit/go-rest-db-goose/infrastructure"
	"github.com/manujdixit/go-rest-db-goose/usecases"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize DB connection
	db, err := sql.Open("mysql", "root:your_password@tcp(localhost:3306)/itemsdb")
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	defer db.Close()

	// Set up repository
	itemRepo := infrastructure.NewMySQLItemRepository(db)

	// Set up use case
	itemUseCase := usecases.NewItemUseCase(itemRepo)

	// Set up controller
	itemController := controllers.NewItemController(itemUseCase)

	// Set up HTTP handlers
	http.HandleFunc("/items", itemController.GetAllItems)
	http.HandleFunc("/items/", itemController.GetItem)
	http.HandleFunc("/items/create", itemController.CreateItem)
	http.HandleFunc("/items/update", itemController.UpdateItem)
	http.HandleFunc("/items/delete", itemController.DeleteItem)

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
