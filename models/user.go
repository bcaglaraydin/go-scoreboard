package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      uuid.UUID `json:"user_id" gorm:"primaryKey"`
	DisplayName string    `json:"display_name" gorm:"type:string;not null; deafult:null"`
	Points      int       `json:"points" gorm:"type:int;not null; deafult:0"`
	Rank        int       `json:"rank" gorm:"type:int; deafult:null"`
	Country     string    `json:"country" gorm:"type:string;not null; deafult:null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = 0
	u.CreatedAt = time.Time{}
	u.UpdatedAt = time.Time{}
	u.DeletedAt = gorm.DeletedAt{}
	return
}
