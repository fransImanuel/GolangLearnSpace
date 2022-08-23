package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fransImanuel/meetmeup/graph/generated"
	"github.com/fransImanuel/meetmeup/graph/model"
)

// CreateMeetUp is the resolver for the createMeetUp field.
func (r *mutationResolver) CreateMeetUp(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
