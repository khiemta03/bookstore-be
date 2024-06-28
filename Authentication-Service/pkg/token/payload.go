package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Contain the payload data of a token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Valid implements jwt.Claims.
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// Return new payload with specific user id and duration
func NewPayload(userId string, duration time.Duration) (*Payload, error) {
	payloadID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        payloadID,
		UserId:    userId,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}
