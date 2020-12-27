package utils

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/Bendomey/goutilities/pkg/validatetoken"
	"github.com/Bendomey/task-assignment/services"
	"github.com/dgrijalva/jwt-go"
)

//UserFromToken unmarshals cliams from jwt to get user id and type
type UserFromToken struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func extractToken(unattendedToken string) (*string, error) {
	if strings.TrimSpace(unattendedToken) == "" {
		return nil, errors.New("AuthorizationFailed")
	}

	//remove bearer
	strArr := strings.Split(unattendedToken, " ")
	if len(strArr) != 2 {
		return nil, errors.New("AuthorizationFailed")
	}

	token := strArr[1]

	return &token, nil
}

// ValidateUser checks if there is a token and then checks if the user is in the db
func ValidateUser(ctx context.Context, userService services.UserService) (*UserFromToken, error) {
	gc, ginErr := GinContextFromContext(ctx)

	if ginErr != nil {
		return nil, ginErr
	}
	//retrieve token from cookie
	unattendedToken := gc.GetHeader("Authorization")

	//extract token
	token, extractTokenErr := extractToken(unattendedToken)
	if extractTokenErr != nil {
		return nil, extractTokenErr
	}

	//extract token metadata
	rawToken, validateError := validatetoken.ValidateJWTToken(*token, os.Getenv("SECRET"))

	if validateError != nil {
		return nil, validateError
	}

	claims, ok := rawToken.Claims.(jwt.MapClaims)
	var userFromTokenImplementation UserFromToken
	if ok && rawToken.Valid {
		userFromTokenImplementation.ID = int(claims["id"].(float64))
		userFromTokenImplementation.Type = claims["type"].(string)
	}

	//check if its exists in db
	_, err := userService.GetUser(ctx, userFromTokenImplementation.ID)
	if err != nil {
		log.Print("from here")
		return nil, err
	}

	return &userFromTokenImplementation, nil
}
