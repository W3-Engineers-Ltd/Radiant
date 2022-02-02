package context

import (
	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

// radiantInput operates the http request header, data, cookie and body.
// it also contains router params and current session.
type radiantInput context.radiantInput

// NewInput return radiantInput generated by Context.
func NewInput() *radiantInput {
	return (*radiantInput)(context.NewInput())
}

// Reset init the radiantInput
func (input *radiantInput) Reset(ctx *Context) {
	(*context.radiantInput)(input).Reset((*context.Context)(ctx))
}

// Protocol returns request protocol name, such as HTTP/1.1 .
func (input *radiantInput) Protocol() string {
	return (*context.radiantInput)(input).Protocol()
}

// URI returns full request url with query string, fragment.
func (input *radiantInput) URI() string {
	return input.Context.Request.RequestURI
}

// URL returns request url path (without query string, fragment).
func (input *radiantInput) URL() string {
	return (*context.radiantInput)(input).URL()
}

// Site returns base site url as scheme://domain type.
func (input *radiantInput) Site() string {
	return (*context.radiantInput)(input).Site()
}

// Scheme returns request scheme as "http" or "https".
func (input *radiantInput) Scheme() string {
	return (*context.radiantInput)(input).Scheme()
}

// Domain returns host name.
// Alias of Host method.
func (input *radiantInput) Domain() string {
	return (*context.radiantInput)(input).Domain()
}

// Host returns host name.
// if no host info in request, return localhost.
func (input *radiantInput) Host() string {
	return (*context.radiantInput)(input).Host()
}

// Method returns http request method.
func (input *radiantInput) Method() string {
	return (*context.radiantInput)(input).Method()
}

// Is returns boolean of this request is on given method, such as Is("POST").
func (input *radiantInput) Is(method string) bool {
	return (*context.radiantInput)(input).Is(method)
}

// IsGet Is this a GET method request?
func (input *radiantInput) IsGet() bool {
	return (*context.radiantInput)(input).IsGet()
}

// IsPost Is this a POST method request?
func (input *radiantInput) IsPost() bool {
	return (*context.radiantInput)(input).IsPost()
}

// IsHead Is this a Head method request?
func (input *radiantInput) IsHead() bool {
	return (*context.radiantInput)(input).IsHead()
}

// IsOptions Is this a OPTIONS method request?
func (input *radiantInput) IsOptions() bool {
	return (*context.radiantInput)(input).IsOptions()
}

// IsPut Is this a PUT method request?
func (input *radiantInput) IsPut() bool {
	return (*context.radiantInput)(input).IsPut()
}

// IsDelete Is this a DELETE method request?
func (input *radiantInput) IsDelete() bool {
	return (*context.radiantInput)(input).IsDelete()
}

// IsPatch Is this a PATCH method request?
func (input *radiantInput) IsPatch() bool {
	return (*context.radiantInput)(input).IsPatch()
}

// IsAjax returns boolean of this request is generated by ajax.
func (input *radiantInput) IsAjax() bool {
	return (*context.radiantInput)(input).IsAjax()
}

// IsSecure returns boolean of this request is in https.
func (input *radiantInput) IsSecure() bool {
	return (*context.radiantInput)(input).IsSecure()
}

// IsWebsocket returns boolean of this request is in webSocket.
func (input *radiantInput) IsWebsocket() bool {
	return (*context.radiantInput)(input).IsWebsocket()
}

// IsUpload returns boolean of whether file uploads in this request or not..
func (input *radiantInput) IsUpload() bool {
	return (*context.radiantInput)(input).IsUpload()
}

// AcceptsHTML Checks if request accepts html response
func (input *radiantInput) AcceptsHTML() bool {
	return (*context.radiantInput)(input).AcceptsHTML()
}

// AcceptsXML Checks if request accepts xml response
func (input *radiantInput) AcceptsXML() bool {
	return (*context.radiantInput)(input).AcceptsXML()
}

// AcceptsJSON Checks if request accepts json response
func (input *radiantInput) AcceptsJSON() bool {
	return (*context.radiantInput)(input).AcceptsJSON()
}

// AcceptsYAML Checks if request accepts json response
func (input *radiantInput) AcceptsYAML() bool {
	return (*context.radiantInput)(input).AcceptsYAML()
}

// IP returns request client ip.
// if in proxy, return first proxy id.
// if error, return RemoteAddr.
func (input *radiantInput) IP() string {
	return (*context.radiantInput)(input).IP()
}

// Proxy returns proxy client ips slice.
func (input *radiantInput) Proxy() []string {
	return (*context.radiantInput)(input).Proxy()
}

// Referer returns http referer header.
func (input *radiantInput) Referer() string {
	return (*context.radiantInput)(input).Referer()
}

// Refer returns http referer header.
func (input *radiantInput) Refer() string {
	return (*context.radiantInput)(input).Refer()
}

// SubDomains returns sub domain string.
// if aa.bb.domain.com, returns aa.bb .
func (input *radiantInput) SubDomains() string {
	return (*context.radiantInput)(input).SubDomains()
}

// Port returns request client port.
// when error or empty, return 80.
func (input *radiantInput) Port() int {
	return (*context.radiantInput)(input).Port()
}

// UserAgent returns request client user agent string.
func (input *radiantInput) UserAgent() string {
	return (*context.radiantInput)(input).UserAgent()
}

// ParamsLen return the length of the params
func (input *radiantInput) ParamsLen() int {
	return (*context.radiantInput)(input).ParamsLen()
}

// Param returns router param by a given key.
func (input *radiantInput) Param(key string) string {
	return (*context.radiantInput)(input).Param(key)
}

// Params returns the map[key]value.
func (input *radiantInput) Params() map[string]string {
	return (*context.radiantInput)(input).Params()
}

// SetParam will set the param with key and value
func (input *radiantInput) SetParam(key, val string) {
	(*context.radiantInput)(input).SetParam(key, val)
}

// ResetParams clears any of the input's Params
// This function is used to clear parameters so they may be reset between filter
// passes.
func (input *radiantInput) ResetParams() {
	(*context.radiantInput)(input).ResetParams()
}

// Query returns input data item string by a given string.
func (input *radiantInput) Query(key string) string {
	return (*context.radiantInput)(input).Query(key)
}

// Header returns request header item string by a given string.
// if non-existed, return empty string.
func (input *radiantInput) Header(key string) string {
	return (*context.radiantInput)(input).Header(key)
}

// Cookie returns request cookie item string by a given key.
// if non-existed, return empty string.
func (input *radiantInput) Cookie(key string) string {
	return (*context.radiantInput)(input).Cookie(key)
}

// Session returns current session item value by a given key.
// if non-existed, return nil.
func (input *radiantInput) Session(key interface{}) interface{} {
	return (*context.radiantInput)(input).Session(key)
}

// CopyBody returns the raw request body data as bytes.
func (input *radiantInput) CopyBody(MaxMemory int64) []byte {
	return (*context.radiantInput)(input).CopyBody(MaxMemory)
}

// Data return the implicit data in the input
func (input *radiantInput) Data() map[interface{}]interface{} {
	return (*context.radiantInput)(input).Data()
}

// GetData returns the stored data in this context.
func (input *radiantInput) GetData(key interface{}) interface{} {
	return (*context.radiantInput)(input).GetData(key)
}

// SetData stores data with given key in this context.
// This data are only available in this context.
func (input *radiantInput) SetData(key, val interface{}) {
	(*context.radiantInput)(input).SetData(key, val)
}

// ParseFormOrMulitForm parseForm or parseMultiForm based on Content-type
func (input *radiantInput) ParseFormOrMulitForm(maxMemory int64) error {
	return (*context.radiantInput)(input).ParseFormOrMultiForm(maxMemory)
}

// Bind data from request.Form[key] to dest
// like /?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie
// var id int  radiantInput.Bind(&id, "id")  id ==123
// var isok bool  radiantInput.Bind(&isok, "isok")  isok ==true
// var ft float64  radiantInput.Bind(&ft, "ft")  ft ==1.2
// ol := make([]int, 0, 2)  radiantInput.Bind(&ol, "ol")  ol ==[1 2]
// ul := make([]string, 0, 2)  radiantInput.Bind(&ul, "ul")  ul ==[str array]
// user struct{Name}  radiantInput.Bind(&user, "user")  user == {Name:"astaxie"}
func (input *radiantInput) Bind(dest interface{}, key string) error {
	return (*context.radiantInput)(input).Bind(dest, key)
}
