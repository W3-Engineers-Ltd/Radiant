package captcha

import (
	"io"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/captcha"
)

// Image struct
type Image captcha.Image

// NewImage returns a new captcha image of the given width and height with the
// given digits, where each digit must be in range 0-9.
func NewImage(digits []byte, width, height int) *Image {
	return (*Image)(captcha.NewImage(digits, width, height))
}

// WriteTo writes captcha image in PNG format into the given writer.
func (m *Image) WriteTo(w io.Writer) (int64, error) {
	return (*captcha.Image)(m).WriteTo(w)
}
