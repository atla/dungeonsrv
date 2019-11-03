package server

import (
	"net/http"

	"github.com/atla/dungeonsrv/pkg/service"
	"github.com/gorilla/mux"
)

//ItemTemplateHandler is the public item handler interface
type ItemTemplateHandler interface {
	GetItemTemplates(w http.ResponseWriter, r *http.Request)
	GetItemTemplateByTemplateID(w http.ResponseWriter, r *http.Request)
}

type itemTemplateHandler struct {
	itemTemplatesService service.ItemTemplatesService
	httpResponder        HTTPResponder
}

//NewItemTemplatesHandler creates a new item handler
func NewItemTemplatesHandler(its service.ItemTemplatesService, httpResponder HTTPResponder) ItemTemplateHandler {
	return &itemTemplateHandler{
		itemTemplatesService: its,
		httpResponder:        httpResponder,
	}
}

// returns the list of item templates
func (ih *itemTemplateHandler) GetItemTemplates(w http.ResponseWriter, r *http.Request) {
	//TODO: check if there is a search/filter
	if items, err := ih.itemTemplatesService.GetItemTemplatesRepository().FindAll(); err != nil {
		ih.httpResponder.ERROR(w, http.StatusNotFound)
	} else {
		ih.httpResponder.JSON(w, http.StatusOK, items)
	}
}

// returns a single item template
func (ih *itemTemplateHandler) GetItemTemplateByTemplateID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var templateID = params["templateID"]

	if itemTemplate, err := ih.itemTemplatesService.GetItemTemplatesRepository().FindByTemplateID(templateID); err != nil {
		ih.httpResponder.ERROR(w, http.StatusNotFound)
	} else {
		ih.httpResponder.JSON(w, http.StatusOK, itemTemplate)
	}
}
