package storage

import (
	"telegram_back/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type storage struct {
	cfg config.Config
	db  *mongo.Database
}

func New(cfg config.Config, db *mongo.Database) *storage {
	return &storage{
		cfg: cfg,
		db:  db,
	}
}
