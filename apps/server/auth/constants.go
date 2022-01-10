package auth

const (
	// Route hit with a GET by Google to complete the "Login With Google"
	// authentication flow.
	loginWithGoogleAuthCallbackRoute = "/api/auth/google/callback"
	// Route his by the user when they want to "Login with Google".
	startLoginWithGoogleRoute = "/api/login/google"
)
