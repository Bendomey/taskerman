package services

import (
	"context"
	"log"
	"time"

	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"github.com/Bendomey/task-assignment/repository"
)

// UserService inteface holds the user-databse transactions of this controller
type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string) (*User, error)
	// UpdateUser(ctx context.Context, name string, phone string, email string, id string) (*User, error)
	// GetUsers(ctx context.Context, skip uint64, take uint64) ([]User, error)
	// GetUser(ctx context.Context, id string) (*User, error)
	// DeleteUser(ctx context.Context, id string) (bool, error)
}

// User model
type User struct {
	ID        int64     `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Type      string    `json:"user_type"`
	IsDeleted bool      `json:"deleted"`
	CreatedBy int64     `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//userRepository gets repository
type userRepository struct {
	repository repository.Repository
}

// NewUserService exposed the repository to the user functions in the module
func NewUserService(r repository.Repository) UserService {
	return &userRepository{r}
}

//save user details here
func (s *userRepository) CreateUser(ctx context.Context, name string, email string, password string) (*User, error) {
	hash, hashErr := hashpassword.HashPassword(password)
	if hashErr != nil {
		return nil, hashErr
	}
	u, err := s.repository.Insert(ctx, "insert into users(fullname,email,password) values($1,$2,$3)", name, email, hash)
	if err != nil {
		return nil, err
	}

	log.Println(u)
	var user = &User{}
	return user, nil
}
