package repository

import "github.com/atla/dungeonsrv/pkg/entities"

//ItemsRepository repository interface
type ItemsRepository interface {
	FindAll() ([]*entities.Item, error)
	FindByID(id string) (*entities.Item, error)
	Store(item *entities.Item) (*entities.Item, error)
	Update(item *entities.Item) error
}
