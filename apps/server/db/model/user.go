package model

import (
	"gorm.io/gorm"
)

// Database model representing an individual user.
type User struct {
	gorm.Model

	// URL to a picture of the user.
	//
	// If nil, no such URL exists.
	AvatarURL *string `gorm:"unique"`
	// Unique email address of the user.
	Email string `gorm:"unique"`
	// First name of the user.
	FirstName string
	// Google identity information of this user.
	Google ThirdPartyIdentity `gorm:"embedded;embeddedPrefix:google_"`
	// Last name of the user.
	LastName string
}

// Column names of the fields belonging to User.
var UserColumnNames = userColumnNamesType{
	AvatarURL: "avatar_url",
	Email:     "email",
	FirstName: "first_name",
	Google: thirdPartyIdentityColumnNames{
		AccessToken:       "google_access_token",
		AccessTokenSecret: "google_access_token_secret",
		Email:             "google_email",
		ExpiresAt:         "google_expires_at",
		IDToken:           "google_id_token",
		UserID:            "google_user_id",
		RefreshToken:      "google_refresh_token",
	},
	LastName:         "last_name",
	modelColumnNames: ModelColumnNames,
}

// Column names of the fields belonging to User.
type userColumnNamesType struct {
	modelColumnNames

	// Column name of the AvatarURL field.
	AvatarURL string
	// Column name of the Email field.
	Email string
	// Column name of the FirstName field.
	FirstName string
	// Column name of the Google field.
	Google thirdPartyIdentityColumnNames
	// Column name of the LastName field.
	LastName string
}
