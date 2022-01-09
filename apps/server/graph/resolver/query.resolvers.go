package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skeswa/chorro/apps/server/graph/model"
	server1 "github.com/skeswa/chorro/apps/server/graph/server"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns server1.QueryResolver implementation.
func (r *Resolver) Query() server1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
