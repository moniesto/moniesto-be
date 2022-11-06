package core

import "github.com/google/uuid"

func CreateID() string {
	id := uuid.New()

	return id.String()
}
