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

//FilterQuery type to help generate filter for queries
type FilterQuery struct {
	Limit     string `json:"limit"`
	Skip      string `json:"skip"`
	Order     string `json:"order"`
	OrderBy   string `json:"orderBy"`
	Search    string `json:"search"`
	DateRange string `json:"dateRange"`
}
