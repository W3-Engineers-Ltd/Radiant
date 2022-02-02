package utils

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

// Display print the data in console
func Display(data ...interface{}) {
	utils.Display(data...)
}

// GetDisplayString return data print string
func GetDisplayString(data ...interface{}) string {
	return utils.GetDisplayString(data...)
}

// Stack get stack bytes
func Stack(skip int, indent string) []byte {
	return utils.Stack(skip, indent)
}
