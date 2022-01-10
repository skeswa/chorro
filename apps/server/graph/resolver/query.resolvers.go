package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/skeswa/chorro/apps/server/graph/model"
	"github.com/skeswa/chorro/apps/server/graph/server"
	"github.com/skeswa/chorro/apps/server/session"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	session := session.ExtractFrom(ctx)

	if !session.IsUserLoggedIn {
		return nil, nil
	}

	user, err := r.DB.User(session.UserID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Email:     &user.Email,
		FirstName: user.FirstName,
		ID:        strconv.FormatUint(uint64(user.ID), 10),
		LastName:  user.LastName,
	}, nil
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
