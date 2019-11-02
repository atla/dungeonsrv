package service

import (
	e "github.com/atla/dungeonsrv/pkg/entities"
	r "github.com/atla/dungeonsrv/pkg/repository"
)

//ItemsService delives logical functions on top of the Items Repo
type ItemsService interface {
	CreateItemFromTemplateID(id string) *e.Item
	GetItemsRepository() r.ItemsRepository
}

type itemsService struct {
	repo r.ItemsRepository
}

//NewItemsService creates a nwe item service
func NewItemsService(itemsRepository r.ItemsRepository) ItemsService {
	return &itemsService{
		repo: itemsRepository,
	}
}

func (is *itemsService) CreateItemFromTemplateID(id string) *e.Item {

	return nil
}

func (is *itemsService) GetItemsRepository() r.ItemsRepository {
	return is.repo
}
