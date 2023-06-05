package core

import (
	"strings"

	"github.com/google/uuid"
)

func CreateID() string {
	id := uuid.New()

	return id.String()
}

func CreatePlainID() string {
	id := uuid.New()

	new_id := strings.ReplaceAll(id.String(), "-", "")

	return new_id
}
