// Package auth provides handlers to enable basic auth support.
// Simple Usage:
//	import(
//		"github.com/W3-Engineers-Ltd/Radiant"
//		"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/auth"
//	)
//
//	func main(){
//		// authenticate every request
//		radiant.InsertFilter("*", radiant.BeforeRouter,auth.Basic("username","secretpassword"))
//		radiant.Run()
//	}
//
//
// Advanced Usage:
//
//	func SecretAuth(username, password string) bool {
//		return username == "astaxie" && password == "helloradiant"
//	}
//	authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
//	radiant.InsertFilter("*", radiant.BeforeRouter,authPlugin)
package auth

import (
	"net/http"

	radiant "github.com/W3-Engineers-Ltd/Radiant/adapter"
	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	beecontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/auth"
)

// Basic is the http basic auth
func Basic(username string, password string) radiant.FilterFunc {
	return func(c *context.Context) {
		f := auth.Basic(username, password)
		f((*beecontext.Context)(c))
	}
}

// NewBasicAuthenticator return the BasicAuth
func NewBasicAuthenticator(secrets SecretProvider, realm string) radiant.FilterFunc {
	f := auth.NewBasicAuthenticator(auth.SecretProvider(secrets), realm)
	return func(c *context.Context) {
		f((*beecontext.Context)(c))
	}
}

// SecretProvider is the SecretProvider function
type SecretProvider auth.SecretProvider

// BasicAuth store the SecretProvider and Realm
type BasicAuth auth.BasicAuth

// CheckAuth Checks the username/password combination from the request. Returns
// either an empty string (authentication failed) or the name of the
// authenticated user.
// Supports MD5 and SHA1 password entries
func (a *BasicAuth) CheckAuth(r *http.Request) string {
	return (*auth.BasicAuth)(a).CheckAuth(r)
}

// RequireAuth http.Handler for BasicAuth which initiates the authentication process
// (or requires reauthentication).
func (a *BasicAuth) RequireAuth(w http.ResponseWriter, r *http.Request) {
	(*auth.BasicAuth)(a).RequireAuth(w, r)
}
