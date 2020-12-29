package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Bendomey/task-assignment/graph/generated"
	"github.com/Bendomey/task-assignment/graph/model"
	"github.com/Bendomey/task-assignment/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	//if there is a validation errorm return the error,else go on with whatever you are doing
	adminData, validateErr := utils.ValidateUser(ctx, r.userService)
	if validateErr != nil {
		return nil, validateErr
	}

	//make sure it is a super admin creating the account
	if adminData.Type != "ADMIN" {
		return nil, errors.New("PermissionDenied")
	}

	// var user model.User
	res, err := r.userService.CreateUser(ctx, input.Name, input.Email, input.Password, string(input.Type), adminData.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginUserInput) (*model.LoginResult, error) {
	res, err := r.userService.LoginUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	return &model.LoginResult{
		User:  &res.User,
		Token: res.Token,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context, filter *model.GetUsersInput, pagination *model.Pagination) ([]*model.User, error) {
	//if there is a validation errorm return the error,else go on with whatever you are doing
	adminData, validateErr := utils.ValidateUser(ctx, r.userService)
	if validateErr != nil {
		return nil, errors.New("AuthorizationFailed")
	}

	//make sure it is a super admin creating the account
	if adminData.Type != "ADMIN" {
		return nil, errors.New("PermissionDenied")
	}

	//generate sieve
	generateQuery, err := utils.GenerateQuery(filter, pagination)
	if err != nil {
		return nil, err
	}

	//if user is sieving with
	userType := ""
	if filter.UserType != nil {
		userType = fmt.Sprintf("AND USER1.user_type='%s'", filter.UserType)
	}

	res, err := r.userService.GetUsers(ctx, *generateQuery, userType)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UsersLength(ctx context.Context, filter *model.GetUsersInput) (int, error) {
	//if there is a validation errorm return the error,else go on with whatever you are doing
	adminData, validateErr := utils.ValidateUser(ctx, r.userService)
	if validateErr != nil {
		return 0, errors.New("AuthorizationFailed")
	}

	//make sure it is a super admin creating the account
	if adminData.Type != "ADMIN" {
		return 0, errors.New("PermissionDenied")
	}

	//generate sieve
	generateQuery, err := utils.GenerateQuery(filter, &model.Pagination{})
	if err != nil {
		return 0, err
	}

	//if user is sieving with
	userType := ""
	if filter.UserType != nil {
		userType = fmt.Sprintf("AND USER1.user_type='%s'", filter.UserType)
	}

	res, err := r.userService.GetUsersLength(ctx, *generateQuery, userType)
	if err != nil {
		return 0, err
	}
	return *res, nil

}

func (r *queryResolver) User(ctx context.Context, filter model.GetUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *Resolver) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: r,
	})
}
