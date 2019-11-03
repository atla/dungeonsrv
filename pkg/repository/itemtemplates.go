package repository

import (
	"io/ioutil"

	"github.com/atla/dungeonsrv/pkg/entities"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//ItemTemplatesRepository repository interface
type ItemTemplatesRepository interface {
	FindAll() ([]*entities.Item, error)
	FindByTemplateID(templateID string) (*entities.Item, error)
}

const (
	itemsSubfolder = "/items"
)

type itemTemplatesRepository struct {
	dungeonDataDir string
}

//NewItemTemplatesRepository creates a new itemTemplatesRepository
func NewItemTemplatesRepository(dungeonDataDir string) ItemTemplatesRepository {
	itr := &itemTemplatesRepository{
		dungeonDataDir: dungeonDataDir,
	}

	weapons := itr.loadItemTemplates("weapons.yaml")
	armor := itr.loadItemTemplates("armor.yaml")

	// load all item templates into memory
	allItemTemplates := []*entities.ItemTemplate{}
	allItemTemplates = append(allItemTemplates, weapons...)
	allItemTemplates = append(allItemTemplates, armor...)

	return itr
}

func (itr *itemTemplatesRepository) loadItemTemplates(yamlFile string) []*entities.ItemTemplate {

	itemTemplates := []*entities.ItemTemplate{}

	fileNameWithPath := itr.dungeonDataDir + yamlFile
	if result, err := ioutil.ReadFile(fileNameWithPath); err != nil {
		log.Error("Could not read item template file ", yamlFile)
	} else {
		if err := yaml.Unmarshal(result, itemTemplates); err != nil {
			log.Error("Error unmarshelling item templates", yamlFile)
		}
	}

	return itemTemplates
}

func (itr *itemTemplatesRepository) FindAll() ([]*entities.Item, error) {
	return nil, nil
}

func (itr *itemTemplatesRepository) FindByTemplateID(templateID string) (*entities.Item, error) {
	return nil, nil
}
