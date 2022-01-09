package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/skeswa/chorro/apps/server/cache"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/chorro/apps/server/db"
	"github.com/skeswa/chorro/apps/server/session"
)

// Initialize everything authentication related for the server.
func Setup(
	cache *cache.Cache,
	config *config.Config,
	db *db.DB,
	mux *http.ServeMux,
) {
	// Tell Goth to use the session store that we've already configured.
	gothic.Store = cache.SessionStore

	// Support for "Login With Google".
	googleAuthProvider := google.New(
		config.Auth.GoogleOAuth2ClientID,
		config.Auth.GoogleOAuth2ClientSecret,
		// Fully-qualified Google auth callback endpoint URL:
		config.PublicBaseUrl()+loginWithGoogleAuthCallbackRoute,
		// Google OAuth 2.0 scopes:
		"email", "profile",
	)

	// Setup Google 3rd-party auth via the Goth authentication library.
	goth.UseProviders(googleAuthProvider)

	// Only allow google auth for now.
	gothic.GetProviderName = func(_ *http.Request) (string, error) {
		return googleAuthProvider.Name(), nil
	}

	// Where we redirect if authentication fails for any reason.
	homeUrlWithAuthFailureFlag := fmt.Sprintf("%s?auth=failed", config.HomeUrl)

	// Starts the "Login With Google" flow if the user isn't already logged in.
	mux.HandleFunc("/login", func(responseWriter http.ResponseWriter, request *http.Request) {
		// GET only.
		if request.Method != http.MethodGet {
			responseWriter.WriteHeader(http.StatusNotFound)

			return
		}

		session := session.Read(cache, request, responseWriter)

		// Go back home if the user is already logged in.
		if couldSkipAuth(
			config,
			db,
			request,
			responseWriter,
			session,
		) {
			return
		}

		gothic.BeginAuthHandler(responseWriter, request)
	})

	// Finishes the "Login With Google" flow.
	mux.HandleFunc(loginWithGoogleAuthCallbackRoute, func(responseWriter http.ResponseWriter, request *http.Request) {
		session := session.Read(cache, request, responseWriter)

		// Go back home if the user is already logged in.
		if couldSkipAuth(
			config,
			db,
			request,
			responseWriter,
			session,
		) {
			return
		}

		googleUser, err := gothic.CompleteUserAuth(responseWriter, request)
		if err != nil {
			log.Println("Failed to complete \"Login With Google\":", err)

			http.Redirect(
				responseWriter,
				request,
				homeUrlWithAuthFailureFlag,
				http.StatusFound,
			)

			return
		}

		if userIDOfGoogleUser, err := db.UserIDOfGoogleUser(
			googleUser.UserID,
		); err != nil {
			log.Printf(
				"Failed to find Google User \"%s\" to complete "+
					"\"Login With Google\": %v\n",
				googleUser.UserID,
				err,
			)

			http.Redirect(
				responseWriter,
				request,
				homeUrlWithAuthFailureFlag,
				http.StatusFound,
			)

			return
		} else {
			if userIDOfGoogleUser == nil {
				// This is an unrecognized Google user, so we should register them.
				if userID, err := db.RegisterGoogleUser(googleUser); err != nil {
					log.Printf(
						"Failed to register new Google User \"%s\" to complete "+
							"\"Login With Google\": %v\n",
						googleUser.UserID,
						err,
					)

					http.Redirect(
						responseWriter,
						request,
						homeUrlWithAuthFailureFlag,
						http.StatusFound,
					)

					return
				} else {
					userIDOfGoogleUser = &userID
				}
			}

			if err := session.LogIn(*userIDOfGoogleUser); err != nil {
				log.Printf(
					"Failed to login Google User \"%s\" to complete "+
						"\"Login With Google\": %v\n",
					googleUser.UserID,
					err,
				)

				http.Redirect(
					responseWriter,
					request,
					homeUrlWithAuthFailureFlag,
					http.StatusFound,
				)

				return
			}
		}

		// Looks like the user is now logged in - let's send them home.
		http.Redirect(
			responseWriter,
			request,
			config.HomeUrl,
			http.StatusFound,
		)
	})
}

// Attempts the short-circuit auth if the user is already logged in by
// redirecting them home.
//
// This function returns true if short circuiting succeeded.
func couldSkipAuth(
	config *config.Config,
	db *db.DB,
	request *http.Request,
	responseWriter http.ResponseWriter,
	session *session.Session) bool {
	// Go back home if the user is already logged in.
	if session.IsUserLoggedIn {
		if doesUserExist, err := db.DoesUserExist(session.UserID); doesUserExist && err == nil {
			http.Redirect(
				responseWriter,
				request,
				config.HomeUrl,
				http.StatusFound,
			)

			return true
		} else if !doesUserExist {
			log.Println("Tried to short-circuit login, but user did not exist")
		} else {
			log.Println(
				"Tried to short-circuit login, but there was an error:",
				err,
			)
		}
	}

	return false
}
