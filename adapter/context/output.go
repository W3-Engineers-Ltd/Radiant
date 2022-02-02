package context

import (
	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

// radiantOutput does work for sending response header.
type radiantOutput context.radiantOutput

// NewOutput returns new radiantOutput.
// it contains nothing now.
func NewOutput() *radiantOutput {
	return (*radiantOutput)(context.NewOutput())
}

// Reset init radiantOutput
func (output *radiantOutput) Reset(ctx *Context) {
	(*context.radiantOutput)(output).Reset((*context.Context)(ctx))
}

// Header sets response header item string via given key.
func (output *radiantOutput) Header(key, val string) {
	(*context.radiantOutput)(output).Header(key, val)
}

// Body sets response body content.
// if EnableGzip, compress content string.
// it sends out response body directly.
func (output *radiantOutput) Body(content []byte) error {
	return (*context.radiantOutput)(output).Body(content)
}

// Cookie sets cookie value via given key.
// others are ordered as cookie's max age time, path,domain, secure and httponly.
func (output *radiantOutput) Cookie(name string, value string, others ...interface{}) {
	(*context.radiantOutput)(output).Cookie(name, value, others)
}

// JSON writes json to response body.
// if encoding is true, it converts utf-8 to \u0000 type.
func (output *radiantOutput) JSON(data interface{}, hasIndent bool, encoding bool) error {
	return (*context.radiantOutput)(output).JSON(data, hasIndent, encoding)
}

// YAML writes yaml to response body.
func (output *radiantOutput) YAML(data interface{}) error {
	return (*context.radiantOutput)(output).YAML(data)
}

// JSONP writes jsonp to response body.
func (output *radiantOutput) JSONP(data interface{}, hasIndent bool) error {
	return (*context.radiantOutput)(output).JSONP(data, hasIndent)
}

// XML writes xml string to response body.
func (output *radiantOutput) XML(data interface{}, hasIndent bool) error {
	return (*context.radiantOutput)(output).XML(data, hasIndent)
}

// ServeFormatted serve YAML, XML OR JSON, depending on the value of the Accept header
func (output *radiantOutput) ServeFormatted(data interface{}, hasIndent bool, hasEncode ...bool) {
	(*context.radiantOutput)(output).ServeFormatted(data, hasIndent, hasEncode...)
}

// Download forces response for download file.
// it prepares the download response header automatically.
func (output *radiantOutput) Download(file string, filename ...string) {
	(*context.radiantOutput)(output).Download(file, filename...)
}

// ContentType sets the content type from ext string.
// MIME type is given in mime package.
func (output *radiantOutput) ContentType(ext string) {
	(*context.radiantOutput)(output).ContentType(ext)
}

// SetStatus sets response status code.
// It writes response header directly.
func (output *radiantOutput) SetStatus(status int) {
	(*context.radiantOutput)(output).SetStatus(status)
}

// IsCachable returns boolean of this request is cached.
// HTTP 304 means cached.
func (output *radiantOutput) IsCachable() bool {
	return (*context.radiantOutput)(output).IsCachable()
}

// IsEmpty returns boolean of this request is empty.
// HTTP 201，204 and 304 means empty.
func (output *radiantOutput) IsEmpty() bool {
	return (*context.radiantOutput)(output).IsEmpty()
}

// IsOk returns boolean of this request runs well.
// HTTP 200 means ok.
func (output *radiantOutput) IsOk() bool {
	return (*context.radiantOutput)(output).IsOk()
}

// IsSuccessful returns boolean of this request runs successfully.
// HTTP 2xx means ok.
func (output *radiantOutput) IsSuccessful() bool {
	return (*context.radiantOutput)(output).IsSuccessful()
}

// IsRedirect returns boolean of this request is redirection header.
// HTTP 301,302,307 means redirection.
func (output *radiantOutput) IsRedirect() bool {
	return (*context.radiantOutput)(output).IsRedirect()
}

// IsForbidden returns boolean of this request is forbidden.
// HTTP 403 means forbidden.
func (output *radiantOutput) IsForbidden() bool {
	return (*context.radiantOutput)(output).IsForbidden()
}

// IsNotFound returns boolean of this request is not found.
// HTTP 404 means not found.
func (output *radiantOutput) IsNotFound() bool {
	return (*context.radiantOutput)(output).IsNotFound()
}

// IsClientError returns boolean of this request client sends error data.
// HTTP 4xx means client error.
func (output *radiantOutput) IsClientError() bool {
	return (*context.radiantOutput)(output).IsClientError()
}

// IsServerError returns boolean of this server handler errors.
// HTTP 5xx means server internal error.
func (output *radiantOutput) IsServerError() bool {
	return (*context.radiantOutput)(output).IsServerError()
}

// Session sets session item value with given key.
func (output *radiantOutput) Session(name interface{}, value interface{}) {
	(*context.radiantOutput)(output).Session(name, value)
}
