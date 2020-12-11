package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/Bendomey/task-assignment/services"

	"github.com/Bendomey/task-assignment/repository"
)

//Resolver get the services
type Resolver struct {
	userService services.UserService
}

// NewGraphqlServer creates a graphql server of all microservices
func NewGraphqlServer(repo repository.Repository) (*Resolver, error) {
	// connect to user service
	userService := services.NewUserService(repo)
	return &Resolver{
		userService: userService,
	}, nil
}
