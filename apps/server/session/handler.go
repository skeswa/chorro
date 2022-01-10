package session

import (
	"net/http"

	"github.com/skeswa/chorro/apps/server/cache"
)

// Middleware that inejcts the Session into the request context of the provided
// handler.
func Handler(
	cache *cache.Cache,
	handler http.Handler,
) http.Handler {
	return http.HandlerFunc(func(
		responseWriter http.ResponseWriter,
		request *http.Request,
	) {
		session := read(cache, request, responseWriter)

		// Embed session into the request context so that is can be conveniently
		// read by handler's functions.
		sessionizedContext := session.embedInto(request.Context())

		// Pass along the sessionized request to handler, so that it can serve its
		// response.
		requestWithSessionizedContext := request.WithContext(sessionizedContext)
		handler.ServeHTTP(responseWriter, requestWithSessionizedContext)
	})
}
