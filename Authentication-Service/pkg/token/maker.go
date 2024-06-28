package token

import "time"

type Maker interface {
	// Create new token
	CreateToken(userId string, duration time.Duration) (string, *Payload, error)

	// Check if the token is valid or not
	ValidateToken(token string) (*Payload, error)
}
