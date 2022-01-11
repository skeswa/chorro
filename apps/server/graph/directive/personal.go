package directive

import (
	"context"
	"errors"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skeswa/chorro/apps/server/graph/model"
	"github.com/skeswa/chorro/apps/server/session"
)

// TODO(skeswa): we should check a generic "ownedBy" field instead.

// Directive used to restrict fields to the currently authenticated user only.
func Personal(context context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if user, isUser := obj.(*model.User); isUser {
		session := session.ExtractFrom(context)
		if !session.IsUserLoggedIn {
			return nil, errors.New("queried field is protected")
		}

		if user.ID != strconv.FormatUint(uint64(session.UserID), 10) {
			return nil, errors.New("queried field is personal")
		}
	}

	return next(context)
}
