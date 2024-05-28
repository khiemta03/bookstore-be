// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type USER struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	FullName       string    `json:"full_name"`
	Address        string    `json:"address"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}
