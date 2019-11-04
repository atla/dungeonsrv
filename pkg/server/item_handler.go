package server

import (
	"encoding/json"
	"net/http"

	"github.com/atla/dungeonsrv/pkg/entities"
	"github.com/atla/dungeonsrv/pkg/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//ItemHandler is the public item handler interface
type ItemHandler interface {
	GetItems(w http.ResponseWriter, r *http.Request)
	GetItemByID(w http.ResponseWriter, r *http.Request)
	PostItem(w http.ResponseWriter, r *http.Request)
	CreateItemFromTemplateID(w http.ResponseWriter, r *http.Request)
}

type itemHandler struct {
	itemsService  service.ItemsService
	httpResponder HTTPResponder
}

//NewItemHandler creates a new item handler
func NewItemHandler(is service.ItemsService, httpResponder HTTPResponder) ItemHandler {
	return &itemHandler{
		itemsService:  is,
		httpResponder: httpResponder,
	}
}

// returns the list of item templates
func (ih *itemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	//TODO: check if there is a search/filter
	if items, err := ih.itemsService.GetItemsRepository().FindAll(); err != nil {
		ih.httpResponder.ERROR(w, http.StatusNotFound)
	} else {
		ih.httpResponder.JSON(w, http.StatusOK, items)
	}
}

// returns a single item
func (ih *itemHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id = params["id"]

	if item, err := ih.itemsService.GetItemsRepository().FindByID(id); err != nil {
		ih.httpResponder.ERROR(w, http.StatusNotFound)
	} else {
		ih.httpResponder.JSON(w, http.StatusOK, item)
	}
}

// Create item from template id
func (ih *itemHandler) CreateItemFromTemplateID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id = params["templateID"]

	item, err := ih.itemsService.CreateItemFromTemplateID(id)

	if err != nil {
		log.Error(err)
		ih.httpResponder.ERROR(w, http.StatusNotFound)
	} else {

		// save any item we created using this method
		ih.itemsService.GetItemsRepository().Store(item)
		ih.httpResponder.JSON(w, http.StatusOK, item)
	}
}

// returns a single item
func (ih *itemHandler) PostItem(w http.ResponseWriter, r *http.Request) {

	var item entities.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	log.WithField("item", item.Name).Info("Creating new Item")

	if storedItem, err := ih.itemsService.GetItemsRepository().Store(&item); err != nil {
		ih.httpResponder.ERROR(w, http.StatusNotFound)
	} else {
		ih.httpResponder.JSON(w, http.StatusOK, storedItem)
	}

}
