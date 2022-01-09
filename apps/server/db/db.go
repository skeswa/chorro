package db

import (
	"github.com/pkg/errors"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/chorro/apps/server/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Interface between the server and its primary data store.
type DB struct {
	// Gorm database client.
	client *gorm.DB
}

// Creates and initializes a new DB.
func New(config *config.Config) (*DB, error) {
	client, err := gorm.Open(
		postgres.Open(config.ForGormPostgresDriver()),
		config.ForGorm(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to database")
	}

	// Initialize the client.
	client.AutoMigrate(&model.User{})

	return &DB{client}, nil
}
