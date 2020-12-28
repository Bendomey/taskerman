package models

import (
	"time"
)

//User model
type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Type      string    `json:"user_type"`
	IsDeleted bool      `json:"deleted"`
	CreatedBy *User     `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
