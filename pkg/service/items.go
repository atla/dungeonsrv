package service

import (
	"errors"

	e "github.com/atla/dungeonsrv/pkg/entities"
	"github.com/atla/dungeonsrv/pkg/generators"
	r "github.com/atla/dungeonsrv/pkg/repository"
	log "github.com/sirupsen/logrus"
)

//ItemsService delives logical functions on top of the Items Repo
type ItemsService interface {
	CreateItemFromTemplateID(id string) (*e.Item, error)
	GetItemsRepository() r.ItemsRepository
}

type itemsService struct {
	repo              r.ItemsRepository
	itemTemplatesRepo r.ItemTemplatesRepository
	generator         generators.ItemGenerator
}

//NewItemsService creates a nwe item service
func NewItemsService(itemsRepository r.ItemsRepository, itemTemplatesRepository r.ItemTemplatesRepository) ItemsService {
	return &itemsService{
		repo:              itemsRepository,
		itemTemplatesRepo: itemTemplatesRepository,
		generator:         generators.NewItemGenerator(itemsRepository, itemTemplatesRepository),
	}
}

func (is *itemsService) CreateItemFromTemplateID(id string) (*e.Item, error) {

	log.Debug("Find item ", id)
	if template, err := is.itemTemplatesRepo.FindByTemplateID(id); err == nil {
		return is.generator.CreateItemWithTemplate(template), nil
	}

	return nil, errors.New("Could not create item from templateID")
}

func (is *itemsService) GetItemsRepository() r.ItemsRepository {
	return is.repo
}
