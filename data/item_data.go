package data

import "github.com/manujdixit/go-rest-db-goose/entities"

type ItemData interface {
	GetAll() ([]entities.Item, error)
	GetByID(id int) (entities.Item, error)
	Create(item entities.Item) (int, error)
	Update(item entities.Item) error
	Delete(id int) error
}
