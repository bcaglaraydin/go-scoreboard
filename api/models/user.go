package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID      uuid.UUID `json:"user_id" `
	DisplayName string    `json:"display_name" `
	Points      int       `json:"points" `
	Rank        int       `json:"rank" `
	Country     string    `json:"country"`
}
