package services

import (
	"context"
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

//User model
type User struct {
	ID        int         `json:"id"`
	Fullname  string      `json:"fullname"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Type      string      `json:"user_type"`
	IsDeleted bool        `json:"deleted"`
	CreatedBy interface{} `json:"createdBy"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
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
	var idRes int
	var createdByRes interface{}
	var nameRes, emailRes, userTypeRes string
	var createdAtRes, updatedAtRes time.Time
	err := s.repository.Insert(context.Background(), "insert into users (fullname,email,password) values($1,$2,$3) returning id,fullname,email,user_type,created_by,created_at,updated_at;", name, email, hash).Scan(
		&idRes, &nameRes, &emailRes, &userTypeRes, &createdByRes, &createdAtRes, &updatedAtRes,
	)

	if err != nil {
		return nil, err
	}

	var user = &User{
		ID:        idRes,
		Fullname:  nameRes,
		Email:     emailRes,
		Type:      userTypeRes,
		CreatedBy: createdByRes,
		CreatedAt: createdAtRes,
		UpdatedAt: updatedAtRes,
	}
	return user, nil
}
