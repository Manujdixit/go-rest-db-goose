// usecases/item_usecase.go
package usecases

import (
	"github.com/manujdixit/go-rest-db-goose/entities"
	"github.com/manujdixit/go-rest-db-goose/repositories"
)

type ItemUseCase struct {
	repo repositories.ItemRepository
}

func NewItemUseCase(repo repositories.ItemRepository) *ItemUseCase {
	return &ItemUseCase{repo: repo}
}

func (uc *ItemUseCase) GetAllItems() ([]entities.Item, error) {
	return uc.repo.GetAll()
}

func (uc *ItemUseCase) GetItemByID(id int) (entities.Item, error) {
	return uc.repo.GetByID(id)
}

func (uc *ItemUseCase) CreateItem(item entities.Item) (int, error) {
	return uc.repo.Create(item)
}

func (uc *ItemUseCase) UpdateItem(item entities.Item) error {
	return uc.repo.Update(item)
}

func (uc *ItemUseCase) DeleteItem(id int) error {
	return uc.repo.Delete(id)
}
