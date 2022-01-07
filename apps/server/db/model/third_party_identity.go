package model

import (
	"time"
)

// Relates a particular user to a 3rd-party authentication provider.
//
// This struct is designed to be embedded within other database models.
type ThirdPartyIdentity struct {
	// String that an OAuth client can use to make requests of the 3rd-party
	// authentication provider.
	//
	// For more, see https://oauth.net/2/access-tokens/.
	AccessToken *string
	// String that sort of acts as the "password" that should be used in
	// conjunction with AccessToken.
	//
	// For more, see https://stackoverflow.com/questions/28057430/what-is-the-access-token-vs-access-token-secret-and-consumer-key-vs-consumer-s.
	AccessTokenSecret *string
	// User email associated with the 3rd-party authentication provider.
	Email *string
	// When AccessToken expires.
	ExpiresAt *time.Time
	// JWT-like string that encodes all of the user's authentication metadata for
	// the 3rd-party authentication provider.
	//
	// For more, see https://www.oauth.com/oauth2-servers/openid-connect/id-tokens/.
	IDToken *string
	// Token that can be used to refresh AccessToken.
	//
	// For more, see https://auth0.com/blog/refresh-tokens-what-are-they-and-when-to-use-them/.
	RefreshToken *string
	// Uniquely identifies the User with the 3rd-party authentication provider.
	UserID *string `gorm:"unique"`
}

// Column names of the fields belonging to ThirdPartyIdentity.
type thirdPartyIdentityColumnNames struct {
	// Column name of the AccessToken field.
	AccessToken string
	// Column name of the AccessTokenSecret field.
	AccessTokenSecret string
	// Column name of the Email field.
	Email string
	// Column name of the ExpiresAt field.
	ExpiresAt string
	// Column name of the IDToken field.
	IDToken string
	// Column name of the RefreshToken field.
	RefreshToken string
	// Column name of the UserID field.
	UserID string
}
