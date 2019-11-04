package service

import (
	r "github.com/atla/dungeonsrv/pkg/repository"
)

//ItemTemplatesService delives logical functions on top of the Items Repo
type ItemTemplatesService interface {
	GetItemTemplatesRepository() r.ItemTemplatesRepository
}

type itemTemplatesService struct {
	repo r.ItemTemplatesRepository
}

//NewItemTemplatesService creates a new ItemTemplatesService
func NewItemTemplatesService(repo r.ItemTemplatesRepository) ItemTemplatesService {
	return &itemTemplatesService{
		repo: repo,
	}
}

func (i *itemTemplatesService) GetItemTemplatesRepository() r.ItemTemplatesRepository {
	return i.repo
}
