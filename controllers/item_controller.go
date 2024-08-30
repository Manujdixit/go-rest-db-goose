package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/manujdixit/go-rest-db-goose/entities"
	"github.com/manujdixit/go-rest-db-goose/usecases"
)

type ItemController struct {
	itemUseCase *usecases.ItemUseCase
}

func NewItemController(uc *usecases.ItemUseCase) *ItemController {
	return &ItemController{itemUseCase: uc}
}

func (c *ItemController) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := c.itemUseCase.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (c *ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := c.itemUseCase.GetItemByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (c *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem entities.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.itemUseCase.CreateItem(newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newItem.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

func (c *ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedItem entities.Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedItem.ID = id
	if err := c.itemUseCase.UpdateItem(updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
}

func (c *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.itemUseCase.DeleteItem(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
