package repository

import (
	"log"

	"my_podcast_api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	*gorm.DB
}

func (mdb *DB) Open(dialect string, details string) (*DB, error) {

	db, err := gorm.Open(dialect, details)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	return &DB{db}, nil
}
