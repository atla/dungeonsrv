package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database is a public interface to a Relational Database using gorm
type Database interface {
	Migrate(object interface{})
	MigrateAll(objects ...interface{})
	Create(value ...interface{})
	Update(value interface{})
	Delete(value interface{})
	Find(out interface{}, where ...interface{}) error
}

/*
func getItemByID(c *gin.Context) {

	id := c.Params.ByName("id")

	var item Item
	if err := db.Find(&item, id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, item)
	}
}

func getItems(c *gin.Context) {
	var items []Item
	if err := db.Find(&items).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, items)
	}
}
*/

// Database ... keeps all necessary db data
type db struct {
	db *gorm.DB
}

// NewDatabase creates a new Database instance and opens with default
func NewDatabase() Database {

	instance := &db{}

	user := "postgres"
	password := "docker"
	host := "localhost"
	databaseName := "dungeonsrv"
	databasePort := "5432"
	var error error
	instance.db, error = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, databasePort, user, databaseName, password))

	//defer instance.db.Close()

	if error != nil {
		log.Fatal(error)
	}

	return instance
}

// Migrate single object
func (db *db) Migrate(object interface{}) {
	db.db.AutoMigrate(object)
}

// MigrateAll all objects
func (db *db) MigrateAll(objects ...interface{}) {
	for _, item := range objects {
		db.Migrate(item)
	}
}

// Create value(s)
func (db *db) Create(values ...interface{}) {
	for _, value := range values {
		db.db.Create(&value)
	}
}

// Update value
func (db *db) Update(value interface{}) {
	db.db.Update(&value)
}

// Delete value
func (db *db) Delete(value interface{}) {
	db.db.Delete(&value)
}

// Find value
func (db *db) Find(out interface{}, where ...interface{}) error {
	return db.db.Find(&out, where).Error
}
