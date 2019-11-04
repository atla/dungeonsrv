package generators

import (
	"github.com/atla/dungeonsrv/pkg/entities"
	"github.com/atla/dungeonsrv/pkg/repository"
	log "github.com/sirupsen/logrus"
)

// ItemGenerator is use to create new items based on prepare templates
// the generation using this generator also supports invocations of scripts
type ItemGenerator interface {
	CreateItemWithTemplate(itemTemplate *entities.ItemTemplate) *entities.Item
}

type itemGenerator struct {
	itemRepository          repository.ItemsRepository
	itemTemplatesRepository repository.ItemTemplatesRepository
}

// NewItemGenerator creates a new item generator
func NewItemGenerator(ir repository.ItemsRepository, itr repository.ItemTemplatesRepository) ItemGenerator {
	return &itemGenerator{
		itemRepository:          ir,
		itemTemplatesRepository: itr,
	}
}

func (ig *itemGenerator) CreateItemWithTemplate(template *entities.ItemTemplate) *entities.Item {

	newItem := &entities.Item{
		Name:        template.Name,
		Description: template.Description,
		ItemType:    template.ItemType,
	}

	log.Debug("Executing post onCreate script")

	// TODO invoke script with item and itemtemplate as context

	// find script with pattern  generators templateID_OnCreate

	// build item_OnCreate context

	// execute script

	// retrieve output

	return newItem
}
