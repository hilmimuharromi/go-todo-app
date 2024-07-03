package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

type PayloadUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
