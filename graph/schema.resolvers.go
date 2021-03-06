package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rs/xid"
	"github.com/shirch/graphql/graph/generated"
	"github.com/shirch/graphql/graph/model"
	"github.com/shirch/graphql/internal/auth"
	"github.com/shirch/graphql/internal/pkg/jwt"
	"github.com/shirch/graphql/internal/users"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	link := model.Link{
		Title:   input.Title,
		Address: input.Address,
		ID:      xid.New().String(),
		User:    user,
	}
	err := r.DB.Create(&link).Error
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	hashedPassword, err := users.HashPassword(input.Password)
	user := model.User{
		Name:     input.Username,
		Password: hashedPassword,
		ID:       xid.New().String(),
	}
	err = r.DB.Create(&user).Error
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user model.User
	user.Name = input.Username
	user.Password = input.Password
	correct := users.Authenticate(&user, r.DB)
	if !correct {
		return "", &WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) UpdateLink(ctx context.Context, linkID string, input model.UpdateLinkInput) (*model.Link, error) {
	user, err := users.GetUserById(input.UserID, r.DB)
	if err != nil {
		return nil, err
	}
	updatedLink := model.Link{
		ID:      linkID,
		Title:   input.Title,
		Address: input.Address,
		User:    &user,
	}
	r.DB.Save(&updatedLink)
	return &updatedLink, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, input model.UserInput) (*model.User, error) {
	updatedUser := model.User{
		ID:       userID,
		Name:     input.Username,
		Password: input.Password,
	}
	r.DB.Save(&updatedUser)
	return &updatedUser, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	r.DB.Find(&resultLinks)

	return resultLinks, nil
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
type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
