package repository

import (
	"errors"
	"io/ioutil"

	"github.com/atla/dungeonsrv/pkg/entities"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//ItemTemplatesRepository repository interface
type ItemTemplatesRepository interface {
	FindAll() ([]entities.ItemTemplate, error)
	FindByTemplateID(templateID string) (*entities.ItemTemplate, error)
}

const (
	itemsSubfolder = "/items/"
)

type itemTemplatesRepository struct {
	dungeonDataDir string
	itemTemplates  []entities.ItemTemplate
}

//NewItemTemplatesRepository creates a new itemTemplatesRepository
func NewItemTemplatesRepository(dungeonDataDir string) ItemTemplatesRepository {
	itr := &itemTemplatesRepository{
		dungeonDataDir: dungeonDataDir,
	}

	weapons := itr.loadItemTemplates("weapons.yaml")
	armor := itr.loadItemTemplates("armor.yaml")

	// load all item templates into memory
	allItemTemplates := []entities.ItemTemplate{}
	allItemTemplates = append(allItemTemplates, weapons...)
	allItemTemplates = append(allItemTemplates, armor...)
	itr.itemTemplates = allItemTemplates

	log.Info("itemTemplateRepo got ", len(allItemTemplates), " templates")

	return itr
}

func (itr *itemTemplatesRepository) loadItemTemplates(yamlFile string) []entities.ItemTemplate {

	itemTemplates := &[]entities.ItemTemplate{}

	fileNameWithPath := itr.dungeonDataDir + itemsSubfolder + yamlFile
	if result, err := ioutil.ReadFile(fileNameWithPath); err != nil {
		log.Error("Could not read item template file ", yamlFile)
	} else {

		log.Debug("CONTENT: ", result)

		if err := yaml.Unmarshal(result, itemTemplates); err != nil {
			log.Error("Error unmarshelling item templates", yamlFile)
			log.Info("-------")
		} else {
			log.Debug("SUCCESS")
			return *itemTemplates
		}
	}

	return *itemTemplates
}

func (itr *itemTemplatesRepository) FindAll() ([]entities.ItemTemplate, error) {
	return itr.itemTemplates, nil
}

func (itr *itemTemplatesRepository) FindByTemplateID(templateID string) (*entities.ItemTemplate, error) {

	for _, v := range itr.itemTemplates {
		if v.TemplateID == templateID {
			return &v, nil
		}
	}

	return nil, errors.New("could not find item template with id: " + templateID)
}
