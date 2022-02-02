package utils

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	return utils.RandomCreateBytes(n, alphabets...)
}
