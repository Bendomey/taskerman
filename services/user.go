package services

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"github.com/Bendomey/goutilities/pkg/signjwt"
	"github.com/Bendomey/goutilities/pkg/validatehash"
	"github.com/Bendomey/task-assignment/graph/model"
	"github.com/Bendomey/task-assignment/models"
	"github.com/Bendomey/task-assignment/repository"
	"github.com/dgrijalva/jwt-go"
)

// UserService inteface holds the user-databse transactions of this controller
type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string, userType string, createdBy int) (*model.User, error)
	LoginUser(ctx context.Context, email string, password string) (*loginResult, error)
	UpdateUser(ctx context.Context, id int, changes string) (*model.User, error)
	GetUsers(ctx context.Context, filter models.FilterQuery, userType string) ([]*model.User, error)
	GetUsersLength(ctx context.Context, filter models.FilterQuery, userType string) (*int, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
	ChangeUserPassword(ctx context.Context, id int, oldPassword string, newPassword string) error
}

// loginResult holds the user details and token
type loginResult struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

//userRepository gets repository
type userRepository struct {
	repository repository.Repository
}

// NewUserService exposed the repository to the user functions in the module
func NewUserService(r repository.Repository) UserService {
	return &userRepository{r}
}

//CreateUser saves user details here
func (s *userRepository) CreateUser(ctx context.Context, name string, email string, password string, userType string, createdBy int) (*model.User, error) {
	hash, hashErr := hashpassword.HashPassword(password)
	if hashErr != nil {
		return nil, hashErr
	}
	var u models.User
	var createdByRes int
	err := s.repository.GetSingle(ctx, "insert into users (fullname,email,password,user_type,created_by) values($1,$2,$3,$4,$5) returning id,fullname,email,user_type,created_by,created_at,updated_at;", name, email, hash, userType, createdBy).Scan(
		&u.ID, &u.Fullname, &u.Email, &u.Type, &createdByRes, &u.CreatedAt, &u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	var user = &model.User{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		UserType: model.UserTypeEnum(u.Type),
		CreatedBy: &model.User{
			ID: createdByRes,
		},
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return user, nil
}

//login user
func (s *userRepository) LoginUser(ctx context.Context, email string, password string) (*loginResult, error) {
	//for user
	var u models.User
	var createdBy models.User

	err := s.repository.GetSingle(ctx,
		"SELECT USER1.id, USER1.fullname, USER1.password, USER1.email, USER1.user_type, USER1.deleted, USER1.created_at, USER1.updated_at, USER2.id, USER2.fullname, USER2.password, USER2.email, USER2.user_type, USER2.deleted, USER2.created_at, USER2.updated_at FROM users AS USER1 LEFT JOIN users AS USER2 ON USER1.created_by=USER2.id WHERE USER1.email=$1 AND USER1.deleted=FALSE",
		email,
	).Scan(&u.ID, &u.Fullname, &u.Password, &u.Email, &u.Type, &u.IsDeleted, &u.CreatedAt, &u.UpdatedAt, &createdBy.ID, &createdBy.Fullname, &createdBy.Password, &createdBy.Email, &createdBy.Type, &createdBy.IsDeleted, &createdBy.CreatedAt, &createdBy.UpdatedAt)
	//
	if err != nil {
		return nil, err
	}

	//since email in db, lets validate hash and then send back
	isSame := validatehash.ValidateCipher(password, u.Password)
	if isSame == false {
		return nil, errors.New("Password is incorrect")
	}

	//sign token
	token, signTokenErrr := signjwt.SignJWT(jwt.MapClaims{
		"id":   u.ID,
		"type": u.Type,
	}, os.Getenv("SECRET"))

	if signTokenErrr != nil {
		return nil, signTokenErrr
	}
	loginResultVar := &loginResult{
		User: model.User{
			ID:       u.ID,
			Fullname: u.Fullname,
			Email:    u.Email,
			UserType: model.UserTypeEnum(u.Type),
			CreatedBy: &model.User{
				ID:        createdBy.ID,
				Fullname:  createdBy.Fullname,
				Email:     createdBy.Email,
				UserType:  model.UserTypeEnum(createdBy.Type),
				CreatedAt: createdBy.CreatedAt,
				UpdatedAt: createdBy.UpdatedAt,
			},
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Token: token,
	}
	return loginResultVar, err
}

//GetUser retrieves a single user
func (s *userRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	var u models.User
	var createdBy models.User

	err := s.repository.GetSingle(ctx,
		"SELECT USER1.id, USER1.fullname, USER1.password, USER1.email, USER1.user_type, USER1.deleted, USER1.created_at, USER1.updated_at, USER2.id, USER2.fullname, USER2.password, USER2.email, USER2.user_type, USER2.deleted, USER2.created_at, USER2.updated_at FROM users AS USER1 LEFT JOIN users AS USER2 ON USER1.created_by=USER2.id WHERE USER1.id=$1 AND USER1.deleted=FALSE",
		id,
	).Scan(&u.ID, &u.Fullname, &u.Password, &u.Email, &u.Type, &u.IsDeleted, &u.CreatedAt, &u.UpdatedAt, &createdBy.ID, &createdBy.Fullname, &createdBy.Password, &createdBy.Email, &createdBy.Type, &createdBy.IsDeleted, &createdBy.CreatedAt, &createdBy.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		UserType: model.UserTypeEnum(u.Type),
		CreatedBy: &model.User{
			ID:        createdBy.ID,
			Fullname:  createdBy.Fullname,
			Email:     createdBy.Email,
			UserType:  model.UserTypeEnum(createdBy.Type),
			CreatedAt: createdBy.CreatedAt,
			UpdatedAt: createdBy.UpdatedAt,
		},
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

// GetUsers retrieves data based on what er have
func (s *userRepository) GetUsers(ctx context.Context, filter models.FilterQuery, userType string) ([]*model.User, error) {
	var users []*model.User
	rows, err := s.repository.GetAll(ctx, fmt.Sprintf("SELECT USER1.id, USER1.fullname, USER1.password, USER1.email, USER1.user_type, USER1.deleted, USER1.created_at, USER1.updated_at, USER2.id, USER2.fullname, USER2.password, USER2.email, USER2.user_type, USER2.deleted, USER2.created_at, USER2.updated_at FROM users AS USER1 LEFT JOIN users AS USER2 ON USER1.created_by=USER2.id WHERE%s%s USER1.deleted=FALSE %s ORDER BY USER1.%s %s %s %s", filter.Search, filter.DateRange, userType, filter.OrderBy, filter.Order, filter.Limit, filter.Skip))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u models.User
		var createdBy models.User
		err := rows.Scan(&u.ID, &u.Fullname, &u.Password, &u.Email, &u.Type, &u.IsDeleted, &u.CreatedAt, &u.UpdatedAt, &createdBy.ID, &createdBy.Fullname, &createdBy.Password, &createdBy.Email, &createdBy.Type, &createdBy.IsDeleted, &createdBy.CreatedAt, &createdBy.UpdatedAt)
		if err != nil {
			return nil, err
		}
		newUser := &model.User{
			ID:       u.ID,
			Fullname: u.Fullname,
			Email:    u.Email,
			UserType: model.UserTypeEnum(u.Type),
			CreatedBy: &model.User{
				ID:        createdBy.ID,
				Fullname:  createdBy.Fullname,
				Email:     createdBy.Email,
				UserType:  model.UserTypeEnum(createdBy.Type),
				CreatedAt: createdBy.CreatedAt,
				UpdatedAt: createdBy.UpdatedAt,
			},
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}
		//append to our users
		users = append(users, newUser)
	}
	return users, nil
}

//GetUsersLength retrieves the length depeneding on a filter
func (s *userRepository) GetUsersLength(ctx context.Context, filter models.FilterQuery, userType string) (*int, error) {
	rows, err := s.repository.GetAll(ctx, fmt.Sprintf("SELECT COUNT(*) FROM users AS USER1 WHERE%s%s USER1.deleted=FALSE %s", filter.Search, filter.DateRange, userType))
	if err != nil {
		return nil, err
	}
	var length *int
	for rows.Next() {
		err := rows.Scan(&length)
		if err != nil {
			return nil, err
		}
	}

	return length, nil
}

//DeleteUser changes the isDeleted to true given the id
func (s *userRepository) DeleteUser(ctx context.Context, id int) error {
	err := s.repository.AlterSingleWithoutReturning(ctx, "UPDATE users SET deleted=TRUE WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

//UpdateUser updates user details
func (s *userRepository) UpdateUser(ctx context.Context, id int, changes string) (*model.User, error) {
	var u models.User
	var createdByRes int
	err := s.repository.GetSingle(ctx, fmt.Sprintf("UPDATE users SET updated_at=now()%s WHERE id=$1 returning id,fullname,email,user_type,created_by,created_at,updated_at;", changes), id).Scan(
		&u.ID, &u.Fullname, &u.Email, &u.Type, &createdByRes, &u.CreatedAt, &u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	var user = &model.User{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		UserType: model.UserTypeEnum(u.Type),
		CreatedBy: &model.User{
			ID: createdByRes,
		},
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return user, nil
}

//ChangeUserPassword updates user password
func (s *userRepository) ChangeUserPassword(ctx context.Context, id int, oldPassword string, newPassword string) error {
	var passwordResult string

	//fetch user first
	geterr := s.repository.GetSingle(ctx, "SELECT password from users WHERE id=$1", id).Scan(&passwordResult)
	if geterr != nil {
		return geterr
	}

	//mnake sure hash is equal old password
	isSame := validatehash.ValidateCipher(oldPassword, passwordResult)
	if isSame == false {
		return errors.New("Old Password is incorrect")
	}

	//hash new password and save
	hash, hashErr := hashpassword.HashPassword(newPassword)
	if hashErr != nil {
		return hashErr
	}

	//update here
	err := s.repository.AlterSingleWithoutReturning(ctx, "UPDATE users SET updated_at=now(), password=$1 WHERE id=$2", hash, id)
	if err != nil {
		return err
	}
	return nil
}
