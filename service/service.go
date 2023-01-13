package service

import (
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/storage"
)

// Service serves logic methods/functions
type Service struct {
	config  config.Config
	Store   *db.Store
	Storage storage.Uploader
}

func NewService(store *db.Store, config config.Config, storage storage.Uploader) (*Service, error) {
	service := &Service{
		config:  config,
		Store:   store,
		Storage: storage,
	}

	return service, nil
}
