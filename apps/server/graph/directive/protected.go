package directive

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skeswa/chorro/apps/server/session"
)

// Directive used to restrict fields to authenticated users only.
func Protected(context context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	session := session.ExtractFrom(context)
	if !session.IsUserLoggedIn {
		return nil, errors.New("queried field is protected")
	}

	return next(context)
}
