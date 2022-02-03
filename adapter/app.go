package adapter

import (
	"net/http"

	context2 "github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

// RadicalApp is an application instance
var RadicalApp *App

func init() {
	// create radiant application
	RadicalApp = (*App)(web.RadicalApp)
}

// App defines radiant application with a new PatternServeMux.
type App web.HttpServer

// NewApp returns a new radiant application.
func NewApp() *App {
	return (*App)(web.NewHttpSever())
}

// MiddleWare function for http.Handler
type MiddleWare web.MiddleWare

// Run radiant application.
func (app *App) Run(mws ...MiddleWare) {
	newMws := oldMiddlewareToNew(mws)
	(*web.HttpServer)(app).Run("", newMws...)
}

func oldMiddlewareToNew(mws []MiddleWare) []web.MiddleWare {
	newMws := make([]web.MiddleWare, 0, len(mws))
	for _, old := range mws {
		newMws = append(newMws, (web.MiddleWare)(old))
	}
	return newMws
}

// Router adds a patterned controller handler to RadicalApp.
// it's an alias method of HttpServer.Router.
// usage:
//  simple router
//  radiant.Router("/admin", &admin.UserController{})
//  radiant.Router("/admin/index", &admin.ArticleController{})
//
//  regex router
//
//  radiant.Router("/api/:id([0-9]+)", &controllers.RController{})
//
//  custom rules
//  radiant.Router("/api/list",&RestController{},"*:ListFood")
//  radiant.Router("/api/create",&RestController{},"post:CreateFood")
//  radiant.Router("/api/update",&RestController{},"put:UpdateFood")
//  radiant.Router("/api/delete",&RestController{},"delete:DeleteFood")
func Router(rootpath string, c ControllerInterface, mappingMethods ...string) *App {
	return (*App)(web.Router(rootpath, c, mappingMethods...))
}

// UnregisterFixedRoute unregisters the route with the specified fixedRoute. It is particularly useful
// in web applications that inherit most routes from a base webapp via the underscore
// import, and aim to overwrite only certain paths.
// The method parameter can be empty or "*" for all HTTP methods, or a particular
// method type (e.g. "GET" or "POST") for selective removal.
//
// Usage (replace "GET" with "*" for all methods):
//  radiant.UnregisterFixedRoute("/yourpreviouspath", "GET")
//  radiant.Router("/yourpreviouspath", yourControllerAddress, "get:GetNewPage")
func UnregisterFixedRoute(fixedRoute string, method string) *App {
	return (*App)(web.UnregisterFixedRoute(fixedRoute, method))
}

// Include will generate router file in the router/xxx.go from the controller's comments
// usage:
// radiant.Include(&BankAccount{}, &OrderController{},&RefundController{},&ReceiptController{})
// type BankAccount struct{
//   radiant.Controller
// }
//
// register the function
// func (b *BankAccount)Mapping(){
//  b.Mapping("ShowAccount" , b.ShowAccount)
//  b.Mapping("ModifyAccount", b.ModifyAccount)
// }
//
// //@router /account/:id  [get]
// func (b *BankAccount) ShowAccount(){
//    //logic
// }
//
//
// //@router /account/:id  [post]
// func (b *BankAccount) ModifyAccount(){
//    //logic
// }
//
// the comments @router url methodlist
// url support all the function Router's pattern
// methodlist [get post head put delete options *]
func Include(cList ...ControllerInterface) *App {
	newList := oldToNewCtrlIntfs(cList)
	return (*App)(web.Include(newList...))
}

func oldToNewCtrlIntfs(cList []ControllerInterface) []web.ControllerInterface {
	newList := make([]web.ControllerInterface, 0, len(cList))
	for _, c := range cList {
		newList = append(newList, c)
	}
	return newList
}

// RESTRouter adds a restful controller handler to RadicalApp.
// its' controller implements radiant.ControllerInterface and
// defines a param "pattern/:objectId" to visit each resource.
func RESTRouter(rootpath string, c ControllerInterface) *App {
	return (*App)(web.RESTRouter(rootpath, c))
}

// AutoRouter adds defined controller handler to RadicalApp.
// it's same to HttpServer.AutoRouter.
// if radiant.AddAuto(&MainController{}) and MainController has methods List and Page,
// visit the url /main/list to exec List function or /main/page to exec Page function.
func AutoRouter(c ControllerInterface) *App {
	return (*App)(web.AutoRouter(c))
}

// AutoPrefix adds controller handler to RadicalApp with prefix.
// it's same to HttpServer.AutoRouterWithPrefix.
// if radiant.AutoPrefix("/admin",&MainController{}) and MainController has methods List and Page,
// visit the url /admin/main/list to exec List function or /admin/main/page to exec Page function.
func AutoPrefix(prefix string, c ControllerInterface) *App {
	return (*App)(web.AutoPrefix(prefix, c))
}

// Get used to register router for Get method
// usage:
//    radiant.Get("/", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Get(rootpath string, f FilterFunc) *App {
	return (*App)(web.Get(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Post used to register router for Post method
// usage:
//    radiant.Post("/api", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Post(rootpath string, f FilterFunc) *App {
	return (*App)(web.Post(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Delete used to register router for Delete method
// usage:
//    radiant.Delete("/api", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Delete(rootpath string, f FilterFunc) *App {
	return (*App)(web.Delete(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Put used to register router for Put method
// usage:
//    radiant.Put("/api", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Put(rootpath string, f FilterFunc) *App {
	return (*App)(web.Put(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Head used to register router for Head method
// usage:
//    radiant.Head("/api", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Head(rootpath string, f FilterFunc) *App {
	return (*App)(web.Head(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Options used to register router for Options method
// usage:
//    radiant.Options("/api", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Options(rootpath string, f FilterFunc) *App {
	return (*App)(web.Options(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Patch used to register router for Patch method
// usage:
//    radiant.Patch("/api", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Patch(rootpath string, f FilterFunc) *App {
	return (*App)(web.Patch(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Any used to register router for all methods
// usage:
//    radiant.Any("/api", func(ctx *context.Context){
//          ctx.Output.Body("hello world")
//    })
func Any(rootpath string, f FilterFunc) *App {
	return (*App)(web.Any(rootpath, func(ctx *context.Context) {
		f((*context2.Context)(ctx))
	}))
}

// Handler used to register a Handler router
// usage:
//    radiant.Handler("/api", http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
//          fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//    }))
func Handler(rootpath string, h http.Handler, options ...interface{}) *App {
	return (*App)(web.Handler(rootpath, h, options))
}

// InsertFilter adds a FilterFunc with pattern condition and action constant.
// The pos means action constant including
// radiant.BeforeStatic, radiant.BeforeRouter, radiant.BeforeExec, radiant.AfterExec and radiant.FinishRouter.
// The bool params is for setting the returnOnOutput value (false allows multiple filters to execute)
func InsertFilter(pattern string, pos int, filter FilterFunc, params ...bool) *App {
	opts := oldToNewFilterOpts(params)
	return (*App)(web.InsertFilter(pattern, pos, func(ctx *context.Context) {
		filter((*context2.Context)(ctx))
	}, opts...))
}
