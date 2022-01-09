package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skeswa/chorro/apps/server/graph/model"
	"github.com/skeswa/chorro/apps/server/graph/server"
)

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo returns server.TodoResolver implementation.
func (r *Resolver) Todo() server.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
