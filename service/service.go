package service

import db "github.com/moniesto/moniesto-be/db/sqlc"

// Service serves logic methods/functions
type Service struct {
	Store *db.Store
}

func NewService(store *db.Store) (*Service, error) {
	service := &Service{
		Store: store,
	}

	return service, nil
}
