// Package apiauth provides handlers to enable apiauth support.
//
// Simple Usage:
//	import(
//		"github.com/W3-Engineers-Ltd/Radiant"
//		"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/apiauth"
//	)
//
//	func main(){
//		// apiauth every request
//		radiant.InsertFilter("*", radiant.BeforeRouter,apiauth.APIBaiscAuth("appid","appkey"))
//		radiant.Run()
//	}
//
// Advanced Usage:
//
//	func getAppSecret(appid string) string {
//		// get appsecret by appid
//		// maybe store in configure, maybe in database
//	}
//
//	radiant.InsertFilter("*", radiant.BeforeRouter,apiauth.APISecretAuth(getAppSecret, 360))
//
// Information:
//
// In the request user should include these params in the query
//
// 1. appid
//
//		 appid is assigned to the application
//
// 2. signature
//
//	get the signature use apiauth.Signature()
//
//	when you send to server remember use url.QueryEscape()
//
// 3. timestamp:
//
//       send the request time, the format is yyyy-mm-dd HH:ii:ss
//
package apiauth

import (
	"net/url"

	radiant "github.com/W3-Engineers-Ltd/Radiant/adapter"
	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	radicalcontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/filter/apiauth"
)

// AppIDToAppSecret is used to get appsecret throw appid
type AppIDToAppSecret apiauth.AppIDToAppSecret

// APIBasicAuth use the basic appid/appkey as the AppIdToAppSecret
func APIBasicAuth(appid, appkey string) radiant.FilterFunc {
	f := apiauth.APIBasicAuth(appid, appkey)
	return func(c *context.Context) {
		f((*radicalcontext.Context)(c))
	}
}

// APIBaiscAuth calls APIBasicAuth for previous callers
func APIBaiscAuth(appid, appkey string) radiant.FilterFunc {
	return APIBasicAuth(appid, appkey)
}

// APISecretAuth use AppIdToAppSecret verify and
func APISecretAuth(f AppIDToAppSecret, timeout int) radiant.FilterFunc {
	ft := apiauth.APISecretAuth(apiauth.AppIDToAppSecret(f), timeout)
	return func(ctx *context.Context) {
		ft((*radicalcontext.Context)(ctx))
	}
}

// Signature used to generate signature with the appsecret/method/params/RequestURI
func Signature(appsecret, method string, params url.Values, requestURL string) string {
	return apiauth.Signature(appsecret, method, params, requestURL)
}
