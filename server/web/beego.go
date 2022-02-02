package web

import (
	"os"
	"path/filepath"
	"sync"
)

const (
	// DEV is for develop
	DEV = "dev"
	// PROD is for production
	PROD = "prod"
)

// M is Map shortcut
type M map[string]interface{}

// Hook function to run
type hookfunc func() error

var hooks = make([]hookfunc, 0) // hook function slice to store the hookfunc

// AddAPPStartHook is used to register the hookfunc
// The hookfuncs will run in radiant.Run()
// such as initiating session , starting middleware , building template, starting admin control and so on.
func AddAPPStartHook(hf ...hookfunc) {
	hooks = append(hooks, hf...)
}

// Run radiant application.
// radiant.Run() default run on HttpPort
// radiant.Run("localhost")
// radiant.Run(":8089")
// radiant.Run("127.0.0.1:8089")
func Run(params ...string) {
	if len(params) > 0 && params[0] != "" {
		BeeApp.Run(params[0])
	}
	BeeApp.Run("")
}

// RunWithMiddleWares Run radiant application with middlewares.
func RunWithMiddleWares(addr string, mws ...MiddleWare) {
	BeeApp.Run(addr, mws...)
}

var initHttpOnce sync.Once

// TODO move to module init function
func initBeforeHTTPRun() {
	initHttpOnce.Do(func() {
		// init hooks
		AddAPPStartHook(
			registerMime,
			registerDefaultErrorHandler,
			registerSession,
			registerTemplate,
			registerAdmin,
			registerGzip,
			// registerCommentRouter,
		)

		for _, hk := range hooks {
			if err := hk(); err != nil {
				panic(err)
			}
		}
	})
}

// TestradiantInit is for test package init
func TestradiantInit(ap string) {
	path := filepath.Join(ap, "conf", "app.conf")
	os.Chdir(ap)
	InitradiantBeforeTest(path)
}

// InitradiantBeforeTest is for test package init
func InitradiantBeforeTest(appConfigPath string) {
	if err := LoadAppConfig(appConfigProvider, appConfigPath); err != nil {
		panic(err)
	}
	BConfig.RunMode = "test"
	initBeforeHTTPRun()
}
