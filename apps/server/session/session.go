package session

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/pkg/errors"
)

// Encapsulates a user-session.
type Session struct {
	// True if the user is logged in.
	IsUserLoggedIn bool
	// Unique identifier of the logged in user if the user is logged in.
	UserID uint

	// Fiber session store wrapped by this struct.
	session *session.Session
}

// Attaches a user to this Session.
func (session *Session) LogIn(userID uint) error {
	rawUserID := strconv.FormatUint(uint64(userID), 10)
	session.session.Set(userIDSessionKey, rawUserID)

	if err := session.session.Save(); err != nil {
		return errors.Wrap(err, "Failed to attach user to session")
	}

	session.IsUserLoggedIn = true
	session.UserID = userID

	return nil
}

// Ends this Session, effectively logging the user out.
func (session *Session) LogOut() error {
	if err := session.session.Destroy(); err != nil {
		return errors.Wrap(err, "Failed to end session")
	}

	session.IsUserLoggedIn = false
	session.UserID = 0

	return nil
}

// Wraps a Fiber session for more ergonomic usage in business logic.
func wrapRawSession(rawSession *session.Session) Session {
	rawUserID, hasRawUserID := rawSession.Get(userIDSessionKey).(string)
	if !hasRawUserID {
		return Session{IsUserLoggedIn: false, session: rawSession}
	}

	userID64, err := strconv.ParseUint(rawUserID, 10, 32)
	if err != nil {
		log.Printf("Failed to convert \"%s\" to a uint64: %v\n", rawUserID, err)

		return Session{IsUserLoggedIn: false, session: rawSession}
	}

	return Session{
		IsUserLoggedIn: true,
		UserID:         uint(userID64),
		session:        rawSession,
	}
}
