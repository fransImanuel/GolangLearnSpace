package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fransimanuel/hackernews/graph/generated"
	"github.com/fransimanuel/hackernews/graph/model"
	"github.com/fransimanuel/hackernews/internal/auth"
	"github.com/fransimanuel/hackernews/internal/links"
	"github.com/fransimanuel/hackernews/internal/users"
	"github.com/fransimanuel/hackernews/pkg/jwt"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	// fmt.Println(input)
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	link.User = user
	linkID := link.Save()
	graphqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title:link.Title, Address:link.Address, User:graphqlUser}, nil

	// panic(fmt.Errorf("not implemented"))
	// link := model.Link{
	// 	Address: input.Address,
	// 	Title: input.Title ,
	// 	User:  &model.User{
	// 		Name: "Test",
	// 	},
	// }

	// return &link, nil	
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	// panic(fmt.Errorf("not implemented"))
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token,nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	// panic(fmt.Errorf("not implemented"))
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct{
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	// panic(fmt.Errorf("not implemented"))
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

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	// var links []*model.Link
	// dummyLink := &model.Link{
	// 	Title: "Our Dummy Link",
	// 	Address: "https://address.org",
	// 	User: &model.User{
	// 		Name: "Admin",
	// 	},
	// }
	// links = append(links, dummyLink)
	// return links, nil
	// panic(fmt.Errorf("not implemented"))
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks{
		graphqlUser := &model.User{
			ID:   link.User.ID,
			Name: link.User.Username,
		}
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: graphqlUser})
	}
	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


