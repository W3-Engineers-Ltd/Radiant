// Package authz provides handlers to enable ACL, RBAC, ABAC authorization support.
// Simple Usage:
//	import(
//		"github.com/W3-Engineers-Ltd/Radiant"
//		"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/authz"
//		"github.com/casbin/casbin"
//	)
//
//	func main(){
//		// mediate the access for every request
//		radiant.InsertFilter("*", radiant.BeforeRouter, authz.NewAuthorizer(casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")))
//		radiant.Run()
//	}
//
//
// Advanced Usage:
//
//	func main(){
//		e := casbin.NewEnforcer("authz_model.conf", "")
//		e.AddRoleForUser("alice", "admin")
//		e.AddPolicy(...)
//
//		radiant.InsertFilter("*", radiant.BeforeRouter, authz.NewAuthorizer(e))
//		radiant.Run()
//	}
package authz

import (
	"net/http"

	"github.com/casbin/casbin"

	radiant "github.com/W3-Engineers-Ltd/Radiant/adapter"
	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	radicalcontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/authz"
)

// NewAuthorizer returns the authorizer.
// Use a casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) radiant.FilterFunc {
	f := authz.NewAuthorizer(e)
	return func(context *context.Context) {
		f((*radicalcontext.Context)(context))
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer authz.BasicAuthorizer

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	return (*authz.BasicAuthorizer)(a).GetUserName(r)
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	return (*authz.BasicAuthorizer)(a).CheckPermission(r)
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter) {
	(*authz.BasicAuthorizer)(a).RequirePermission(w)
}
