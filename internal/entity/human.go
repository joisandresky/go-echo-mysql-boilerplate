package entity

import (
	"time"
)

type Human struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Race      string    `json:"race" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
