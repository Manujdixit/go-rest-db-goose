// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manujdixit/go-rest-db-goose/controllers"
	"github.com/manujdixit/go-rest-db-goose/storage"
	"github.com/manujdixit/go-rest-db-goose/usecases"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:your_password@tcp(localhost:3306)/itemsdb")
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	defer db.Close()

	itemRepo := storage.NewMySQLItemRepository(db)
	itemUseCase := usecases.NewItemUseCase(itemRepo)
	itemController := controllers.NewItemController(itemUseCase)

	http.HandleFunc("/items", itemController.GetAllItems)
	http.HandleFunc("/items/", itemController.GetItem)
	http.HandleFunc("/items/create", itemController.CreateItem)
	http.HandleFunc("/items/update", itemController.UpdateItem)
	http.HandleFunc("/items/delete", itemController.DeleteItem)

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
