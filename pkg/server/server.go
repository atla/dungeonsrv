package server

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"net/http"
	"time"

	"github.com/atla/dungeonsrv/pkg/db"

	"github.com/atla/dungeonsrv/pkg/service"
	"github.com/gorilla/mux"
)

// App ... main application structure
type App interface {
	Run()
}

type app struct {

	// generic app base
	Router *mux.Router
	routes Routes
	db     *db.Client

	// dungeonsrv specific
	dungeonDataDir string
	facade         service.Facade
}

// Route ... contains data to declare a route
type Route struct {
	Pattern     string
	Method      string
	Name        string
	HandlerFunc http.HandlerFunc
}

// Routes ... type to hold multiple routes
type Routes []Route

//HTTPResponder serves method to respond to http calls
type HTTPResponder interface {
	JSON(w http.ResponseWriter, code int, payload interface{})
	ERROR(w http.ResponseWriter, code int)
}

// NewApp returns an application instance
// this is the primary stateless server providing an API interface
func NewApp(dungeonDataDir string) App {
	app := &app{
		db:             db.New(),
		dungeonDataDir: dungeonDataDir,
	}

	app.facade = service.NewFacade(app.db, dungeonDataDir)

	return app
}

// JSON responds to the request with the given code and payload
func (app *app) JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

// JSON responds to the request with the given code and payload
func (app *app) ERROR(w http.ResponseWriter, code int) {
	response := []byte("Error")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger function for http calls
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// SetupRoutes ... Configures the routes
func (app *app) setupRoutes() {

	itemHandler := NewItemHandler(app.facade.ItemsService(), app)
	itemTemplateHandler := NewItemTemplatesHandler(app.facade.ItemTemplatesService(), app)

	app.routes = Routes{
		//items
		Route{"/api/items", "GET", "Get all items", itemHandler.GetItems},
		Route{"/api/items/{id}", "GET", "Get a single item", itemHandler.GetItemByID},
		Route{"/api/items", "POST", "Create a new item", itemHandler.PostItem},
		Route{"/api/createItem/{templateID}", "PUT", "Create a new item from template id", itemHandler.CreateItemFromTemplateID},

		//itemTemplates
		Route{"/api/itemTemplates", "GET", "Get all itemTemplates", itemTemplateHandler.GetItemTemplates},
		Route{"/api/itemTemplates/{templateID}", "GET", "Get a single itemTemplate by templateID", itemTemplateHandler.GetItemTemplateByTemplateID},
	}

	// wrap all routes in logger
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range app.routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	app.Router = router

	// also setup static serving
	app.Router.PathPrefix("/app").Handler(http.FileServer(http.Dir("public/")))
}

// Run ... starts the server
func (app *app) Run() {

	app.db.Connect("mongodb://localhost:27017")
	app.setupRoutes()

	fmt.Println("Server is running, listening on port 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", app.Router))
}
