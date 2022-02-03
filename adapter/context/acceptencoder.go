// Copyright 2021 radiant Author. All Rights Reserved.
//

package context

import (
	"io"
	"net/http"
	"os"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

// InitGzip init the gzipcompress
func InitGzip(minLength, compressLevel int, methods []string) {
	context.InitGzip(minLength, compressLevel, methods)
}

// WriteFile reads from file and writes to writer by the specific encoding(gzip/deflate)
func WriteFile(encoding string, writer io.Writer, file *os.File) (bool, string, error) {
	return context.WriteFile(encoding, writer, file)
}

// WriteBody reads  writes content to writer by the specific encoding(gzip/deflate)
func WriteBody(encoding string, writer io.Writer, content []byte) (bool, string, error) {
	return context.WriteBody(encoding, writer, content)
}

// ParseEncoding will extract the right encoding for response
// the Accept-Encoding's sec is here:
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.3
func ParseEncoding(r *http.Request) string {
	return context.ParseEncoding(r)
}
