package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      uuid.UUID `json:"user_id" gorm:"primaryKey"`
	DisplayName string    `json:"display_name" gorm:"type:string;not null; deafult:null"`
	Points      int       `json:"points" gorm:"type:int;not null; deafult:0"`
	Rank        int       `json:"rank" gorm:"type:int; deafult:null"`
}
