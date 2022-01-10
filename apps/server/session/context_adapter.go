package session

import (
	"context"
	"log"
)

// A private key for context that only this package can access.
//
// This is important to prevent collisions between different context uses.
var contextKey = &contextKeyType{sessionName}

// Extracts the Session embedded within the specified context, panicing if no
// such embedded Session exists.
func ExtractFrom(context context.Context) *Session {
	session, hasSession := context.Value(contextKey).(*Session)
	if !hasSession {
		log.Panicf("Context \"%v\" lacks an embedded session", context)
	}

	return session
}

// Unique type that wraps a specified name to create a private key.
type contextKeyType struct {
	// String wrapped by this struct that creates a unique, debuggable, private
	// key.
	name string
}

// Returns a new copy of the specified originalContext that includes this
// Session.
func (session *Session) embedInto(originalContext context.Context) context.Context {
	return context.WithValue(originalContext, contextKey, session)
}
