package session

import (
	"context"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
	"github.com/W3-Engineers-Ltd/Radiant/server/web"
	webContext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/session"
)

// Session maintain session for web service
// Session new a session storage and store it into webContext.Context
// experimental feature, we may change this in the future
func Session(providerType session.ProviderType, options ...session.ManagerConfigOpt) web.FilterChain {
	sessionConfig := session.NewManagerConfig(options...)
	sessionManager, _ := session.NewManager(string(providerType), sessionConfig)
	go sessionManager.GC()

	return func(next web.FilterFunc) web.FilterFunc {
		return func(ctx *webContext.Context) {
			if ctx.Input.CruSession != nil {
				return
			}

			if sess, err := sessionManager.SessionStart(ctx.ResponseWriter, ctx.Request); err != nil {
				logs.Error(`init session error:%s`, err.Error())
			} else {
				// release session at the end of request
				defer sess.SessionRelease(context.Background(), ctx.ResponseWriter)
				ctx.Input.CruSession = sess
			}

			next(ctx)
		}
	}
}
