package service

import (
	"context"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

func (service *Service) GetAllActivePosts() ([]db.GetAllActivePostsRow, error) {

	ctx := context.Background()

	posts, err := service.Store.GetAllActivePosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, err
}
