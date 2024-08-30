package infrastructure

import (
	"database/sql"

	"github.com/manujdixit/go-rest-db-goose/entities"
	"github.com/manujdixit/go-rest-db-goose/repositories"
)

type MySQLItemRepository struct {
	DB *sql.DB
}

func NewMySQLItemRepository(db *sql.DB) repositories.ItemRepository {
	return &MySQLItemRepository{DB: db}
}

func (r *MySQLItemRepository) GetAll() ([]entities.Item, error) {
	rows, err := r.DB.Query("SELECT id, name FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entities.Item
	for rows.Next() {
		var item entities.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *MySQLItemRepository) GetByID(id int) (entities.Item, error) {
	var item entities.Item
	err := r.DB.QueryRow("SELECT id, name FROM items WHERE id = ?", id).Scan(&item.ID, &item.Name)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *MySQLItemRepository) Create(item entities.Item) (int, error) {
	result, err := r.DB.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *MySQLItemRepository) Update(item entities.Item) error {
	_, err := r.DB.Exec("UPDATE items SET name = ? WHERE id = ?", item.Name, item.ID)
	return err
}

func (r *MySQLItemRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM items WHERE id = ?", id)
	return err
}
