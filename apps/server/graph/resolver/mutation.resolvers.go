package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skeswa/chorro/apps/server/graph/model"
	server1 "github.com/skeswa/chorro/apps/server/graph/server"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns server1.MutationResolver implementation.
func (r *Resolver) Mutation() server1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
