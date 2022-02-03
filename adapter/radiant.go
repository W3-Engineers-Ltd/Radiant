package adapter

import (
	"github.com/W3-Engineers-Ltd/Radiant"
	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

const (

	// VERSION represent radiant web framework version.
	VERSION = radiant.VERSION

	// DEV is for develop
	DEV = web.DEV
	// PROD is for production
	PROD = web.PROD
)

// M is Map shortcut
type M web.M

// Hook function to run
type hookfunc func() error

// AddAPPStartHook is used to register the hookfunc
// The hookfuncs will run in radiant.Run()
// such as initiating session , starting middleware , building template, starting admin control and so on.
func AddAPPStartHook(hf ...hookfunc) {
	for i := 0; i < len(hf); i++ {
		f := hf[i]
		web.AddAPPStartHook(func() error {
			return f()
		})
	}
}

// Run radiant application.
// radiant.Run() default run on HttpPort
// radiant.Run("localhost")
// radiant.Run(":8089")
// radiant.Run("127.0.0.1:8089")
func Run(params ...string) {
	web.Run(params...)
}

// RunWithMiddleWares Run radiant application with middlewares.
func RunWithMiddleWares(addr string, mws ...MiddleWare) {
	newMws := oldMiddlewareToNew(mws)
	web.RunWithMiddleWares(addr, newMws...)
}

// TestRadiantInit is for test package init
func TestRadiantInit(ap string) {
	web.TestRadiantInit(ap)
}

// InitRadiantBeforeTest is for test package init
func InitRadiantBeforeTest(appConfigPath string) {
	web.InitRadiantBeforeTest(appConfigPath)
}
