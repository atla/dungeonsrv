package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/atla/dungeonsrv/pkg/db"
	"github.com/atla/dungeonsrv/pkg/repository"

	"github.com/atla/dungeonsrv/pkg/service"
	"github.com/gorilla/mux"
)

// Route ... contains data to declare a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ... type to hold multiple routes
type Routes []Route

// App ... main application structure
type App interface {
	Run()
}

type app struct {
	Router *mux.Router
	routes Routes
	db     *db.Client
}

//HTTPResponder serves method to respond to http calls
type HTTPResponder interface {
	JSON(w http.ResponseWriter, code int, payload interface{})
	ERROR(w http.ResponseWriter, code int)
}

// NewApp returns an application instance
// this is the primary stateless server providing an API interface
func NewApp() App {
	return &app{
		db: db.New(),
	}

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

	itemRepository := repository.NewMongoItemsRepository(app.db)
	itemsService := service.NewItemsService(itemRepository)
	itemHandler := NewItemHandler(itemsService, app)

	app.routes = Routes{
		Route{
			"Get all items",
			"GET",
			"/api/items",
			itemHandler.GetItems,
		},
		Route{
			"Get a single item",
			"GET",
			"/api/items/{id}",
			itemHandler.GetItemByID,
		},
		Route{
			"Create a new item",
			"POST",
			"/api/items",
			itemHandler.PostItem,
		},
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
