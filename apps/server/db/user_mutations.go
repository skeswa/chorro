package db

import (
	"net/mail"
	"net/url"

	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"github.com/skeswa/chorro/apps/server/db/model"
)

// Creates an entry in the database for the specified googleUser.
func (db *DB) RegisterGoogleUser(googleUser goth.User) (uint, error) {
	user := model.User{
		Google: model.ThirdPartyIdentity{
			AccessToken:       &googleUser.AccessToken,
			AccessTokenSecret: &googleUser.AccessTokenSecret,
			Email:             &googleUser.Email,
			ExpiresAt:         &googleUser.ExpiresAt,
			IDToken:           &googleUser.IDToken,
			RefreshToken:      &googleUser.RefreshToken,
		},
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
	}

	if len(googleUser.AvatarURL) > 0 {
		if _, err := url.ParseRequestURI(googleUser.AvatarURL); err != nil {
			errors.Wrap(err, "Failed to register Google user - invalid avatar url")
		} else {
			user.AvatarURL = &googleUser.AvatarURL
		}
	}

	if _, err := mail.ParseAddress(googleUser.Email); err != nil {
		errors.Wrap(err, "Failed to register Google user - invalid email")
	} else {
		user.Email = googleUser.Email
	}

	if len(googleUser.UserID) < 1 {
		return 0, errors.New("Failed to register Google user - invalid user id")
	} else {
		user.Google.UserID = &googleUser.UserID
	}

	if err := db.client.Create(&user).
		Error; err != nil {
		return 0, errors.Wrap(err, "Failed to register Google user - user creation failed")
	}

	return user.ID, nil
}
