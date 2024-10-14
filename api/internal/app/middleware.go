package app

import (
	"fmt"
	"net/http"
)

// RequireAuthentication middleware to protect routes
// RequireAuthentication middleware to protect routes
func (app *App) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Println(" * * * 🧄 RequireAuthentication middleware invoked for:", r.URL.Path)

		// Check if the session store is initialized
		if app.Store == nil {
			app.Logger.Println(" * * * ⛔️ Session store is not initialized")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Check if the request has a valid session
		session, err := app.Store.Get(r, "session")
		if err != nil {
			app.Logger.Println(" * * * ⛔️ Error getting session:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Check if the session contains the userName
		userName, ok := session.Values["userName"]
		if !ok {
			app.Logger.Println(" * * * ⛔️ User name not found in session")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Check if the userName is a string
		userNameStr, ok := userName.(string)
		if !ok {
			app.Logger.Println(" * * * ⛔️ User name is not a string")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Log the authenticated user
		app.Logger.Println(" * * * ✅ User authenticated:", userNameStr)

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// IsAuthenticated checks if the user is authenticated
func (app *App) IsAuthenticated(r *http.Request) (bool, string) {
	fmt.Println("IsAuthenticated invoked")
	session, err := app.Store.Get(r, "session")
	if err != nil {
		fmt.Println("Error getting session:", err)
		app.Logger.Println(" * * * ⛔️ Error getting session:", err)
		return false, ""
	}
	userName, ok := session.Values["userName"]
	if !ok {
		fmt.Println("User name not found in session")
		app.Logger.Println(" * * * ⛔️ User name not found in session")
		return false, ""
	}
	userNameStr, ok := userName.(string)
	if !ok {
		fmt.Println("User name is not a string")
		app.Logger.Println(" * * * ⛔️ User name is not a string")
		return false, ""
	}
	fmt.Println("User authenticated:", userNameStr)
	app.Logger.Println(" * * * ✅ User authenticated:", userNameStr)
	return true, userNameStr
}
