package adapter

import (
	"html/template"
	"io"
	"net/http"

	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

// ExecuteTemplate applies the template with name  to the specified data object,
// writing the output to wr.
// A template will be executed safely in parallel.
func ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return web.ExecuteTemplate(wr, name, data)
}

// ExecuteViewPathTemplate applies the template with name and from specific viewPath to the specified data object,
// writing the output to wr.
// A template will be executed safely in parallel.
func ExecuteViewPathTemplate(wr io.Writer, name string, viewPath string, data interface{}) error {
	return web.ExecuteViewPathTemplate(wr, name, viewPath, data)
}

// AddFuncMap let user to register a func in the template.
func AddFuncMap(key string, fn interface{}) error {
	return web.AddFuncMap(key, fn)
}

type templatePreProcessor func(root, path string, funcs template.FuncMap) (*template.Template, error)

type templateFile struct {
	root  string
	files map[string][]string
}

// HasTemplateExt return this path contains supported template extension of radiant or not.
func HasTemplateExt(paths string) bool {
	return web.HasTemplateExt(paths)
}

// AddTemplateExt add new extension for template.
func AddTemplateExt(ext string) {
	web.AddTemplateExt(ext)
}

// AddViewPath adds a new path to the supported view paths.
// Can later be used by setting a controller ViewPath to this folder
// will panic if called after radiant.Run()
func AddViewPath(viewPath string) error {
	return web.AddViewPath(viewPath)
}

// BuildTemplate will build all template files in a directory.
// it makes radiant can render any template file in view directory.
func BuildTemplate(dir string, files ...string) error {
	return web.BuildTemplate(dir, files...)
}

type templateFSFunc func() http.FileSystem

func defaultFSFunc() http.FileSystem {
	return FileSystem{}
}

// SetTemplateFSFunc set default filesystem function
func SetTemplateFSFunc(fnt templateFSFunc) {
	web.SetTemplateFSFunc(func() http.FileSystem {
		return fnt()
	})
}

// SetViewsPath sets view directory path in radiant application.
func SetViewsPath(path string) *App {
	return (*App)(web.SetViewsPath(path))
}

// SetStaticPath sets static directory path and proper url pattern in radiant application.
// if radiant.SetStaticPath("static","public"), visit /static/* to load static file in folder "public".
func SetStaticPath(url string, path string) *App {
	return (*App)(web.SetStaticPath(url, path))
}

// DelStaticPath removes the static folder setting in this url pattern in radiant application.
func DelStaticPath(url string) *App {
	return (*App)(web.DelStaticPath(url))
}

// AddTemplateEngine add a new templatePreProcessor which support extension
func AddTemplateEngine(extension string, fn templatePreProcessor) *App {
	return (*App)(web.AddTemplateEngine(extension, func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		return fn(root, path, funcs)
	}))
}
