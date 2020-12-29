// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CreateUserInput struct {
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Type     UserTypeEnum `json:"type"`
}

type DateRange struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type GetUsersInput struct {
	UserType     *UserTypeEnum `json:"user_type"`
	Order        *OrderType    `json:"order"`
	OrderBy      *string       `json:"orderBy"`
	DateField    *string       `json:"dateField"`
	DateRange    *DateRange    `json:"dateRange"`
	Search       *string       `json:"search"`
	SearchFields []string      `json:"searchFields"`
}

type LoginResult struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Pagination struct {
	Skip  *int `json:"skip"`
	Limit *int `json:"limit"`
}

type User struct {
	ID        int          `json:"id"`
	Fullname  string       `json:"fullname"`
	Email     string       `json:"email"`
	UserType  UserTypeEnum `json:"user_type"`
	CreatedBy *User        `json:"createdBy"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

type OrderType string

const (
	OrderTypeAsc  OrderType = "ASC"
	OrderTypeDesc OrderType = "DESC"
)

var AllOrderType = []OrderType{
	OrderTypeAsc,
	OrderTypeDesc,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeAsc, OrderTypeDesc:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserTypeEnum string

const (
	UserTypeEnumAdmin    UserTypeEnum = "ADMIN"
	UserTypeEnumUser     UserTypeEnum = "USER"
	UserTypeEnumReviewer UserTypeEnum = "REVIEWER"
)

var AllUserTypeEnum = []UserTypeEnum{
	UserTypeEnumAdmin,
	UserTypeEnumUser,
	UserTypeEnumReviewer,
}

func (e UserTypeEnum) IsValid() bool {
	switch e {
	case UserTypeEnumAdmin, UserTypeEnumUser, UserTypeEnumReviewer:
		return true
	}
	return false
}

func (e UserTypeEnum) String() string {
	return string(e)
}

func (e *UserTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserTypeEnum", str)
	}
	return nil
}

func (e UserTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
