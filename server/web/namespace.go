package web

import (
	"net/http"
	"strings"

	radicalcontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

type namespaceCond func(*radicalcontext.Context) bool

// LinkNamespace used as link action
type LinkNamespace func(*Namespace)

// Namespace is store all the info
type Namespace struct {
	prefix   string
	handlers *ControllerRegister
}

// NewNamespace get new Namespace
func NewNamespace(prefix string, params ...LinkNamespace) *Namespace {
	ns := &Namespace{
		prefix:   prefix,
		handlers: NewControllerRegister(),
	}
	for _, p := range params {
		p(ns)
	}
	return ns
}

// Cond set condition function
// if cond return true can run this namespace, else can't
// usage:
// ns.Cond(func (ctx *context.Context) bool{
//       if ctx.Input.Domain() == "api.radiant.me" {
//         return true
//       }
//       return false
//   })
// Cond as the first filter
func (n *Namespace) Cond(cond namespaceCond) *Namespace {
	fn := func(ctx *radicalcontext.Context) {
		if !cond(ctx) {
			exception("405", ctx)
		}
	}
	if v := n.handlers.filters[BeforeRouter]; len(v) > 0 {
		mr := new(FilterRouter)
		mr.tree = NewTree()
		mr.pattern = "*"
		mr.filterFunc = fn
		mr.tree.AddRouter("*", true)
		n.handlers.filters[BeforeRouter] = append([]*FilterRouter{mr}, v...)
	} else {
		n.handlers.InsertFilter("*", BeforeRouter, fn)
	}
	return n
}

// Filter add filter in the Namespace
// action has before & after
// FilterFunc
// usage:
// Filter("before", func (ctx *context.Context){
//       _, ok := ctx.Input.Session("uid").(int)
//       if !ok && ctx.Request.RequestURI != "/login" {
//          ctx.Redirect(302, "/login")
//        }
//   })
func (n *Namespace) Filter(action string, filter ...FilterFunc) *Namespace {
	var a int
	if action == "before" {
		a = BeforeRouter
	} else if action == "after" {
		a = FinishRouter
	}
	for _, f := range filter {
		n.handlers.InsertFilter("*", a, f, WithReturnOnOutput(true))
	}
	return n
}

// Router same as radiant.Rourer
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Router
func (n *Namespace) Router(rootpath string, c ControllerInterface, mappingMethods ...string) *Namespace {
	n.handlers.Add(rootpath, c, WithRouterMethods(c, mappingMethods...))
	return n
}

// AutoRouter same as radiant.AutoRouter
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#AutoRouter
func (n *Namespace) AutoRouter(c ControllerInterface) *Namespace {
	n.handlers.AddAuto(c)
	return n
}

// AutoPrefix same as radiant.AutoPrefix
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#AutoPrefix
func (n *Namespace) AutoPrefix(prefix string, c ControllerInterface) *Namespace {
	n.handlers.AddAutoPrefix(prefix, c)
	return n
}

// Get same as radiant.Get
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Get
func (n *Namespace) Get(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Get(rootpath, f)
	return n
}

// Post same as radiant.Post
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Post
func (n *Namespace) Post(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Post(rootpath, f)
	return n
}

// Delete same as radiant.Delete
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Delete
func (n *Namespace) Delete(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Delete(rootpath, f)
	return n
}

// Put same as radiant.Put
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Put
func (n *Namespace) Put(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Put(rootpath, f)
	return n
}

// Head same as radiant.Head
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Head
func (n *Namespace) Head(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Head(rootpath, f)
	return n
}

// Options same as radiant.Options
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Options
func (n *Namespace) Options(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Options(rootpath, f)
	return n
}

// Patch same as radiant.Patch
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Patch
func (n *Namespace) Patch(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Patch(rootpath, f)
	return n
}

// Any same as radiant.Any
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Any
func (n *Namespace) Any(rootpath string, f HandleFunc) *Namespace {
	n.handlers.Any(rootpath, f)
	return n
}

// Handler same as radiant.Handler
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Handler
func (n *Namespace) Handler(rootpath string, h http.Handler) *Namespace {
	n.handlers.Handler(rootpath, h)
	return n
}

// Include add include class
// refer: https://godoc.org/github.com/W3-Engineers-Ltd/Radiant#Include
func (n *Namespace) Include(cList ...ControllerInterface) *Namespace {
	n.handlers.Include(cList...)
	return n
}

// CtrlGet same as radiant.CtrlGet
func (n *Namespace) CtrlGet(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlGet(rootpath, f)
	return n
}

// CtrlPost same as radiant.CtrlPost
func (n *Namespace) CtrlPost(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlPost(rootpath, f)
	return n
}

// CtrlDelete same as radiant.CtrlDelete
func (n *Namespace) CtrlDelete(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlDelete(rootpath, f)
	return n
}

// CtrlPut same as radiant.CtrlPut
func (n *Namespace) CtrlPut(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlPut(rootpath, f)
	return n
}

// CtrlHead same as radiant.CtrlHead
func (n *Namespace) CtrlHead(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlHead(rootpath, f)
	return n
}

// CtrlOptions same as radiant.CtrlOptions
func (n *Namespace) CtrlOptions(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlOptions(rootpath, f)
	return n
}

// CtrlPatch same as radiant.CtrlPatch
func (n *Namespace) CtrlPatch(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlPatch(rootpath, f)
	return n
}

// Any same as radiant.CtrlAny
func (n *Namespace) CtrlAny(rootpath string, f interface{}) *Namespace {
	n.handlers.CtrlAny(rootpath, f)
	return n
}

// Namespace add nest Namespace
// usage:
// ns := radiant.NewNamespace(“/v1”).
// Namespace(
//    radiant.NewNamespace("/shop").
//        Get("/:id", func(ctx *context.Context) {
//            ctx.Output.Body([]byte("shopinfo"))
//    }),
//    radiant.NewNamespace("/order").
//        Get("/:id", func(ctx *context.Context) {
//            ctx.Output.Body([]byte("orderinfo"))
//    }),
//    radiant.NewNamespace("/crm").
//        Get("/:id", func(ctx *context.Context) {
//            ctx.Output.Body([]byte("crminfo"))
//    }),
// )
func (n *Namespace) Namespace(ns ...*Namespace) *Namespace {
	for _, ni := range ns {
		for k, v := range ni.handlers.routers {
			if _, ok := n.handlers.routers[k]; ok {
				addPrefix(v, ni.prefix)
				n.handlers.routers[k].AddTree(ni.prefix, v)
			} else {
				t := NewTree()
				t.AddTree(ni.prefix, v)
				addPrefix(t, ni.prefix)
				n.handlers.routers[k] = t
			}
		}
		if ni.handlers.enableFilter {
			for pos, filterList := range ni.handlers.filters {
				for _, mr := range filterList {
					t := NewTree()
					t.AddTree(ni.prefix, mr.tree)
					mr.tree = t
					n.handlers.insertFilterRouter(pos, mr)
				}
			}
		}
	}
	return n
}

// AddNamespace register Namespace into radiant.Handler
// support multi Namespace
func AddNamespace(nl ...*Namespace) {
	for _, n := range nl {
		for k, v := range n.handlers.routers {
			if _, ok := RadicalApp.Handlers.routers[k]; ok {
				addPrefix(v, n.prefix)
				RadicalApp.Handlers.routers[k].AddTree(n.prefix, v)
			} else {
				t := NewTree()
				t.AddTree(n.prefix, v)
				addPrefix(t, n.prefix)
				RadicalApp.Handlers.routers[k] = t
			}
		}
		if n.handlers.enableFilter {
			for pos, filterList := range n.handlers.filters {
				for _, mr := range filterList {
					t := NewTree()
					t.AddTree(n.prefix, mr.tree)
					mr.tree = t
					RadicalApp.Handlers.insertFilterRouter(pos, mr)
				}
			}
		}
	}
}

func addPrefix(t *Tree, prefix string) {
	for _, v := range t.fixrouters {
		addPrefix(v, prefix)
	}
	if t.wildcard != nil {
		addPrefix(t.wildcard, prefix)
	}
	for _, l := range t.leaves {
		if c, ok := l.runObject.(*ControllerInfo); ok {
			if !strings.HasPrefix(c.pattern, prefix) {
				c.pattern = prefix + c.pattern
			}
		}
	}
}

// NSCond is Namespace Condition
func NSCond(cond namespaceCond) LinkNamespace {
	return func(ns *Namespace) {
		ns.Cond(cond)
	}
}

// NSBefore Namespace BeforeRouter filter
func NSBefore(filterList ...FilterFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Filter("before", filterList...)
	}
}

// NSAfter add Namespace FinishRouter filter
func NSAfter(filterList ...FilterFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Filter("after", filterList...)
	}
}

// NSInclude Namespace Include ControllerInterface
func NSInclude(cList ...ControllerInterface) LinkNamespace {
	return func(ns *Namespace) {
		ns.Include(cList...)
	}
}

// NSRouter call Namespace Router
func NSRouter(rootpath string, c ControllerInterface, mappingMethods ...string) LinkNamespace {
	return func(ns *Namespace) {
		ns.Router(rootpath, c, mappingMethods...)
	}
}

// NSGet call Namespace Get
func NSGet(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Get(rootpath, f)
	}
}

// NSPost call Namespace Post
func NSPost(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Post(rootpath, f)
	}
}

// NSHead call Namespace Head
func NSHead(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Head(rootpath, f)
	}
}

// NSPut call Namespace Put
func NSPut(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Put(rootpath, f)
	}
}

// NSDelete call Namespace Delete
func NSDelete(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Delete(rootpath, f)
	}
}

// NSAny call Namespace Any
func NSAny(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Any(rootpath, f)
	}
}

// NSOptions call Namespace Options
func NSOptions(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Options(rootpath, f)
	}
}

// NSPatch call Namespace Patch
func NSPatch(rootpath string, f HandleFunc) LinkNamespace {
	return func(ns *Namespace) {
		ns.Patch(rootpath, f)
	}
}

// NSCtrlGet call Namespace CtrlGet
func NSCtrlGet(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlGet(rootpath, f)
	}
}

// NSCtrlPost call Namespace CtrlPost
func NSCtrlPost(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlPost(rootpath, f)
	}
}

// NSCtrlHead call Namespace CtrlHead
func NSCtrlHead(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlHead(rootpath, f)
	}
}

// NSCtrlPut call Namespace CtrlPut
func NSCtrlPut(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlPut(rootpath, f)
	}
}

// NSCtrlDelete call Namespace CtrlDelete
func NSCtrlDelete(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlDelete(rootpath, f)
	}
}

// NSCtrlAny call Namespace CtrlAny
func NSCtrlAny(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlAny(rootpath, f)
	}
}

// NSCtrlOptions call Namespace CtrlOptions
func NSCtrlOptions(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlOptions(rootpath, f)
	}
}

// NSCtrlPatch call Namespace CtrlPatch
func NSCtrlPatch(rootpath string, f interface{}) LinkNamespace {
	return func(ns *Namespace) {
		ns.CtrlPatch(rootpath, f)
	}
}

// NSAutoRouter call Namespace AutoRouter
func NSAutoRouter(c ControllerInterface) LinkNamespace {
	return func(ns *Namespace) {
		ns.AutoRouter(c)
	}
}

// NSAutoPrefix call Namespace AutoPrefix
func NSAutoPrefix(prefix string, c ControllerInterface) LinkNamespace {
	return func(ns *Namespace) {
		ns.AutoPrefix(prefix, c)
	}
}

// NSNamespace add sub Namespace
func NSNamespace(prefix string, params ...LinkNamespace) LinkNamespace {
	return func(ns *Namespace) {
		n := NewNamespace(prefix, params...)
		ns.Namespace(n)
	}
}

// NSHandler add handler
func NSHandler(rootpath string, h http.Handler) LinkNamespace {
	return func(ns *Namespace) {
		ns.Handler(rootpath, h)
	}
}
