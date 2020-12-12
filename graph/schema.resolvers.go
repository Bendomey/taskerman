package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Bendomey/task-assignment/graph/generated"
	"github.com/Bendomey/task-assignment/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	// var user model.User
	res, err := r.userService.CreateUser(ctx, input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	u := &model.User{
		ID:        res.ID,
		Fullname:  res.Fullname,
		Email:     res.Email,
		CreatedBy: nil,
		UserType:  model.UserTypeEnum(res.Type),
		CreatedAt: res.CreatedAt.String(),
		UpdatedAt: res.UpdatedAt.String(),
	}
	return u, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginUserInput) (*model.LoginResult, error) {
	res, err := r.userService.LoginUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	return &model.LoginResult{
		User: &model.User{
			ID:        res.User.ID,
			Fullname:  res.User.Fullname,
			Email:     res.User.Email,
			CreatedBy: nil,
			UserType:  model.UserTypeEnum(res.User.Type),
			CreatedAt: res.User.CreatedAt.String(),
			UpdatedAt: res.User.UpdatedAt.String(),
		},
		Token: res.Token,
	}, nil

	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
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
