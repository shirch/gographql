package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"github.com/shirch/graphql/graph/generated"
	"github.com/shirch/graphql/graph/model"
	"github.com/shirch/graphql/internal/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.LinkInput) (*model.Link, error) {
	link := model.Link{
		Title:   input.Title,
		Address: input.Address,
		ID:      xid.New().String(),
	}
	err := r.DB.Create(&link).Error
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	hashedPassword, err := HashPassword(input.Password)
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLink(ctx context.Context, linkID string, input model.LinkInput) (*model.Link, error) {
	updatedLink := model.Link{
		ID:      linkID,
		Title:   input.Title,
		Address: input.Address,
	}
	r.DB.Save(&updatedLink)
	return &updatedLink, nil
}

func (r *mutationResolver) DeleteLink(ctx context.Context, linkID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
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
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GetUserIdByUsername(username string, DB *gorm.DB) (string, error) {
	var id string
	DB.Raw("select ID from Users WHERE Username = ?", username).Scan(&id)
	return id, nil
}
