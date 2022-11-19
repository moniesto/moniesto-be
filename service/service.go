package service

import (
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
)

// Service serves logic methods/functions
type Service struct {
	config config.Config
	Store  *db.Store
}

func NewService(store *db.Store, config config.Config) (*Service, error) {
	service := &Service{
		config: config,
		Store:  store,
	}

	return service, nil
}
