package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/moniesto/moniesto-be/util"
)

// Types of the Token errors
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID   `json:"id"`
	User      UserPayload `json:"user"`
	IssuedAt  time.Time   `json:"issued_at"`
	ExpiredAt time.Time   `json:"expired_at"`
}

type GeneralPaylod struct {
	UserPayload UserPayload `json:"user"`
}

type UserPayload struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(generalPaylod GeneralPaylod, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID: tokenID,
		User: UserPayload{
			ID:       generalPaylod.UserPayload.ID,
			Username: generalPaylod.UserPayload.Username,
		},
		IssuedAt:  util.Now(),
		ExpiredAt: util.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if util.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
