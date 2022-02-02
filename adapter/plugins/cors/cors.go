// Package cors provides handlers to enable CORS support.
// Usage
//	import (
// 		"github.com/W3-Engineers-Ltd/Radiant"
//		"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/cors"
// )
//
//	func main() {
//		// CORS for https://foo.* origins, allowing:
//		// - PUT and PATCH methods
//		// - Origin header
//		// - Credentials share
//		radiant.InsertFilter("*", radiant.BeforeRouter, cors.Allow(&cors.Options{
//			AllowOrigins:     []string{"https://*.foo.com"},
//			AllowMethods:     []string{"PUT", "PATCH"},
//			AllowHeaders:     []string{"Origin"},
//			ExposeHeaders:    []string{"Content-Length"},
//			AllowCredentials: true,
//		}))
//		radiant.Run()
//	}
package cors

import (
	radiant "github.com/W3-Engineers-Ltd/Radiant/adapter"
	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	beecontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/cors"
)

// Options represents Access Control options.
type Options cors.Options

// Header converts options into CORS headers.
func (o *Options) Header(origin string) (headers map[string]string) {
	return (*cors.Options)(o).Header(origin)
}

// PreflightHeader converts options into CORS headers for a preflight response.
func (o *Options) PreflightHeader(origin, rMethod, rHeaders string) (headers map[string]string) {
	return (*cors.Options)(o).PreflightHeader(origin, rMethod, rHeaders)
}

// IsOriginAllowed looks up if the origin matches one of the patterns
// generated from Options.AllowOrigins patterns.
func (o *Options) IsOriginAllowed(origin string) bool {
	return (*cors.Options)(o).IsOriginAllowed(origin)
}

// Allow enables CORS for requests those match the provided options.
func Allow(opts *Options) radiant.FilterFunc {
	f := cors.Allow((*cors.Options)(opts))
	return func(c *context.Context) {
		f((*beecontext.Context)(c))
	}
}
