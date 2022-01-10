package session

import (
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"github.com/skeswa/chorro/apps/server/cache"
)

// Encapsulates the session information encoded within a particular request.
type Session struct {
	// True if the user is logged in.
	IsUserLoggedIn bool
	// Unique identifier of the logged in user if the user is logged in.
	UserID uint

	// Session instance wrapped by this struct.
	inner *sessions.Session
	// HTTP request with which this Session is associated.
	request *http.Request
	// HTTP repsonse writer paired with request.
	responseWriter http.ResponseWriter
}

// Attaches a user to this Session.
func (session *Session) LogIn(userID uint) error {
	userIDString := strconv.FormatUint(uint64(userID), 10)
	session.inner.Values[userIDSessionKey] = userIDString

	if err := session.inner.Save(session.request, session.responseWriter); err != nil {
		return errors.Wrapf(err, "Failed to attach user with \"%d\" to session", userID)
	}

	session.IsUserLoggedIn = true
	session.UserID = userID

	return nil
}

// Ends this Session, effectively logging the user out.
func (session *Session) LogOut() error {
	// Expire the session to "destroy" it.
	session.inner.Options.MaxAge = -1

	if err := session.inner.Save(session.request, session.responseWriter); err != nil {
		return errors.Wrap(err, "Failed to destroy the session")
	}

	session.IsUserLoggedIn = false
	session.UserID = 0

	return nil
}

// Reads the Session of the specified request, and its associated response writer.
func read(cache *cache.Cache, request *http.Request, responseWriter http.ResponseWriter) *Session {
	// Get the innersession. We're ignoring the error resulting from decoding an
	// existing session: Get() always returns a session, even if empty.
	inner, _ := cache.SessionStore.Get(request, sessionName)

	session := Session{
		IsUserLoggedIn: false,
		UserID:         0,

		inner:          inner,
		request:        request,
		responseWriter: responseWriter,
	}

	rawUserID, hasRawUserID := inner.Values[userIDSessionKey]
	if hasRawUserID {
		if userIDString, isRawUserIDAString := rawUserID.(string); isRawUserIDAString {
			if userIDInt64, err := strconv.ParseUint(userIDString, 10, 32); err == nil {
				session.IsUserLoggedIn = true
				session.UserID = uint(userIDInt64)
			}
		}
	}

	return &session
}
