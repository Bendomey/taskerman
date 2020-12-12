package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"github.com/Bendomey/goutilities/pkg/signjwt"
	"github.com/Bendomey/goutilities/pkg/validatehash"
	"github.com/Bendomey/task-assignment/repository"
	"github.com/dgrijalva/jwt-go"
)

// UserService inteface holds the user-databse transactions of this controller
type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string) (*User, error)
	LoginUser(ctx context.Context, email string, password string) (*loginResult, error)
	// UpdateUser(ctx context.Context, name string, phone string, email string, id string) (*User, error)
	// GetUsers(ctx context.Context, skip uint64, take uint64) ([]User, error)
	// GetUser(ctx context.Context, id string) (*User, error)
	// DeleteUser(ctx context.Context, id string) (bool, error)
}

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

// loginResult holds the user details and token
type loginResult struct {
	User  User   `json:"user"`
	Token string `json:"token"`
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
	var nameRes, emailRes, userTypeRes string
	var createdAtRes, updatedAtRes time.Time
	err := s.repository.GetSingle(ctx, "insert into users (fullname,email,password) values($1,$2,$3) returning id,fullname,email,user_type,created_at,updated_at;", name, email, hash).Scan(
		&idRes, &nameRes, &emailRes, &userTypeRes, &createdAtRes, &updatedAtRes,
	)

	if err != nil {
		return nil, err
	}

	var user = &User{
		ID:        idRes,
		Fullname:  nameRes,
		Email:     emailRes,
		Type:      userTypeRes,
		CreatedBy: nil,
		CreatedAt: createdAtRes,
		UpdatedAt: updatedAtRes,
	}
	return user, nil
}

//login user
func (s *userRepository) LoginUser(ctx context.Context, email string, password string) (*loginResult, error) {
	var idRes int
	var fullname, emailRes, passwordRes, userTypeRes string
	var createdAtRes, updatedAtRes time.Time

	err := s.repository.GetSingle(ctx, "select id,fullname, email,password,user_type,created_at,updated_at from users where email=$1", email).Scan(&idRes, &fullname, &emailRes, &passwordRes, &userTypeRes, &createdAtRes, &updatedAtRes)
	if err != nil {
		return nil, err
	}
	//since email in db, lets validate hash and then send back
	isSame := validatehash.ValidateCipher(password, passwordRes)
	if isSame == false {
		return nil, errors.New("Password is incorrect")
	}

	//sign token
	token, signTokenErrr := signjwt.SignJWT(jwt.MapClaims{
		"id":   idRes,
		"type": userTypeRes,
	}, os.Getenv("SECRET"))

	if signTokenErrr != nil {
		return nil, signTokenErrr
	}
	loginResultVar := &loginResult{
		User: User{
			ID:        idRes,
			Fullname:  fullname,
			Email:     emailRes,
			Type:      userTypeRes,
			CreatedAt: createdAtRes,
			UpdatedAt: updatedAtRes,
		},
		Token: token,
	}

	return loginResultVar, err
}
