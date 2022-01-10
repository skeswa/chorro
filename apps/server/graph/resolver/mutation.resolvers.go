package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/skeswa/chorro/apps/server/graph/server"
	"github.com/skeswa/chorro/apps/server/session"
)

func (r *mutationResolver) LogOut(ctx context.Context) (bool, error) {
	session := session.ExtractFrom(ctx)

	if err := session.LogOut(); err != nil {
		return false, err
	}

	return true, nil
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
