package service

import (
	"context"
	"fmt"

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

func (service *Service) UpdatePostStatus(activePost db.GetAllActivePostsRow) error {
	fmt.Println("UpdatePostStatus", activePost)
	return nil
}
