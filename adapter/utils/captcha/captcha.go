// Package captcha implements generation and verification of image CAPTCHAs.
// an example for use captcha
//
// ```
// package controllers
//
// import (
// 	"github.com/W3-Engineers-Ltd/Radiant"
// 	"github.com/W3-Engineers-Ltd/Radiant/client/cache"
// 	"github.com/W3-Engineers-Ltd/Radiant/server/web/captcha"
// )
//
// var cpt *captcha.Captcha
//
// func init() {
// 	// use radiant cache system store the captcha data
// 	store := cache.NewMemoryCache()
// 	cpt = captcha.NewWithFilter("/captcha/", store)
// }
//
// type MainController struct {
// 	radiant.Controller
// }
//
// func (this *MainController) Get() {
// 	this.TplName = "index.tpl"
// }
//
// func (this *MainController) Post() {
// 	this.TplName = "index.tpl"
//
// 	this.Data["Success"] = cpt.VerifyReq(this.Ctx.Request)
// }
// ```
//
// template usage
//
// ```
// {{.Success}}
// <form action="/" method="post">
// 	{{create_captcha}}
// 	<input name="captcha" type="text">
// </form>
// ```
package captcha

import (
	"html/template"
	"net/http"
	"time"

	"github.com/W3-Engineers-Ltd/Radiant/adapter/cache"
	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/captcha"
	radicalcontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

var defaultChars = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

const (
	// default captcha attributes
	challengeNums    = 6
	expiration       = 600 * time.Second
	fieldIDName      = "captcha_id"
	fieldCaptchaName = "captcha"
	cachePrefix      = "captcha_"
	defaultURLPrefix = "/captcha/"
)

// Captcha struct
type Captcha captcha.Captcha

// Handler radiant filter handler for serve captcha image
func (c *Captcha) Handler(ctx *context.Context) {
	(*captcha.Captcha)(c).Handler((*radicalcontext.Context)(ctx))
}

// CreateCaptchaHTML template func for output html
func (c *Captcha) CreateCaptchaHTML() template.HTML {
	return (*captcha.Captcha)(c).CreateCaptchaHTML()
}

// CreateCaptcha create a new captcha id
func (c *Captcha) CreateCaptcha() (string, error) {
	return (*captcha.Captcha)(c).CreateCaptcha()
}

// VerifyReq verify from a request
func (c *Captcha) VerifyReq(req *http.Request) bool {
	return (*captcha.Captcha)(c).VerifyReq(req)
}

// Verify direct verify id and challenge string
func (c *Captcha) Verify(id string, challenge string) (success bool) {
	return (*captcha.Captcha)(c).Verify(id, challenge)
}

// NewCaptcha create a new captcha.Captcha
func NewCaptcha(urlPrefix string, store cache.Cache) *Captcha {
	return (*Captcha)(captcha.NewCaptcha(urlPrefix, cache.CreateOldToNewAdapter(store)))
}

// NewWithFilter create a new captcha.Captcha and auto AddFilter for serve captacha image
// and add a template func for output html
func NewWithFilter(urlPrefix string, store cache.Cache) *Captcha {
	return (*Captcha)(captcha.NewWithFilter(urlPrefix, cache.CreateOldToNewAdapter(store)))
}
