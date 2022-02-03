package context

import (
	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

// RadiantOutput does work for sending response header.
type RadiantOutput context.RadiantOutput

// NewOutput returns new RadiantOutput.
// it contains nothing now.
func NewOutput() *RadiantOutput {
	return (*RadiantOutput)(context.NewOutput())
}

// Reset init RadiantOutput
func (output *RadiantOutput) Reset(ctx *Context) {
	(*context.RadiantOutput)(output).Reset((*context.Context)(ctx))
}

// Header sets response header item string via given key.
func (output *RadiantOutput) Header(key, val string) {
	(*context.RadiantOutput)(output).Header(key, val)
}

// Body sets response body content.
// if EnableGzip, compress content string.
// it sends out response body directly.
func (output *RadiantOutput) Body(content []byte) error {
	return (*context.RadiantOutput)(output).Body(content)
}

// Cookie sets cookie value via given key.
// others are ordered as cookie's max age time, path,domain, secure and httponly.
func (output *RadiantOutput) Cookie(name string, value string, others ...interface{}) {
	(*context.RadiantOutput)(output).Cookie(name, value, others)
}

// JSON writes json to response body.
// if encoding is true, it converts utf-8 to \u0000 type.
func (output *RadiantOutput) JSON(data interface{}, hasIndent bool, encoding bool) error {
	return (*context.RadiantOutput)(output).JSON(data, hasIndent, encoding)
}

// YAML writes yaml to response body.
func (output *RadiantOutput) YAML(data interface{}) error {
	return (*context.RadiantOutput)(output).YAML(data)
}

// JSONP writes jsonp to response body.
func (output *RadiantOutput) JSONP(data interface{}, hasIndent bool) error {
	return (*context.RadiantOutput)(output).JSONP(data, hasIndent)
}

// XML writes xml string to response body.
func (output *RadiantOutput) XML(data interface{}, hasIndent bool) error {
	return (*context.RadiantOutput)(output).XML(data, hasIndent)
}

// ServeFormatted serve YAML, XML OR JSON, depending on the value of the Accept header
func (output *RadiantOutput) ServeFormatted(data interface{}, hasIndent bool, hasEncode ...bool) {
	(*context.RadiantOutput)(output).ServeFormatted(data, hasIndent, hasEncode...)
}

// Download forces response for download file.
// it prepares the download response header automatically.
func (output *RadiantOutput) Download(file string, filename ...string) {
	(*context.RadiantOutput)(output).Download(file, filename...)
}

// ContentType sets the content type from ext string.
// MIME type is given in mime package.
func (output *RadiantOutput) ContentType(ext string) {
	(*context.RadiantOutput)(output).ContentType(ext)
}

// SetStatus sets response status code.
// It writes response header directly.
func (output *RadiantOutput) SetStatus(status int) {
	(*context.RadiantOutput)(output).SetStatus(status)
}

// IsCachable returns boolean of this request is cached.
// HTTP 304 means cached.
func (output *RadiantOutput) IsCachable() bool {
	return (*context.RadiantOutput)(output).IsCachable()
}

// IsEmpty returns boolean of this request is empty.
// HTTP 201，204 and 304 means empty.
func (output *RadiantOutput) IsEmpty() bool {
	return (*context.RadiantOutput)(output).IsEmpty()
}

// IsOk returns boolean of this request runs well.
// HTTP 200 means ok.
func (output *RadiantOutput) IsOk() bool {
	return (*context.RadiantOutput)(output).IsOk()
}

// IsSuccessful returns boolean of this request runs successfully.
// HTTP 2xx means ok.
func (output *RadiantOutput) IsSuccessful() bool {
	return (*context.RadiantOutput)(output).IsSuccessful()
}

// IsRedirect returns boolean of this request is redirection header.
// HTTP 301,302,307 means redirection.
func (output *RadiantOutput) IsRedirect() bool {
	return (*context.RadiantOutput)(output).IsRedirect()
}

// IsForbidden returns boolean of this request is forbidden.
// HTTP 403 means forbidden.
func (output *RadiantOutput) IsForbidden() bool {
	return (*context.RadiantOutput)(output).IsForbidden()
}

// IsNotFound returns boolean of this request is not found.
// HTTP 404 means not found.
func (output *RadiantOutput) IsNotFound() bool {
	return (*context.RadiantOutput)(output).IsNotFound()
}

// IsClientError returns boolean of this request client sends error data.
// HTTP 4xx means client error.
func (output *RadiantOutput) IsClientError() bool {
	return (*context.RadiantOutput)(output).IsClientError()
}

// IsServerError returns boolean of this server handler errors.
// HTTP 5xx means server internal error.
func (output *RadiantOutput) IsServerError() bool {
	return (*context.RadiantOutput)(output).IsServerError()
}

// Session sets session item value with given key.
func (output *RadiantOutput) Session(name interface{}, value interface{}) {
	(*context.RadiantOutput)(output).Session(name, value)
}
