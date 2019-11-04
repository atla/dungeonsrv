package service

import (
	"github.com/atla/dungeonsrv/pkg/db"
	"github.com/atla/dungeonsrv/pkg/repository"
)

//Facade ...
type Facade interface {
	ItemsService() ItemsService
	ItemTemplatesService() ItemTemplatesService
}

type facade struct {
	is  ItemsService
	its ItemTemplatesService

	db *db.Client
}

//NewFacade creates a new service facade
func NewFacade(db *db.Client, dungeonDataDir string) Facade {

	itemsRepo := repository.NewMongoItemsRepository(db)
	itemTemplatesRepo := repository.NewItemTemplatesRepository(dungeonDataDir)

	return &facade{
		is:  NewItemsService(itemsRepo, itemTemplatesRepo),
		its: NewItemTemplatesService(itemTemplatesRepo),
	}
}

func (f *facade) ItemsService() ItemsService {
	return f.is
}

func (f *facade) ItemTemplatesService() ItemTemplatesService {
	return f.its
}
