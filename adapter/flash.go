package adapter

import (
	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

// FlashData is a tools to maintain data when using across request.
type FlashData web.FlashData

// NewFlash return a new empty FlashData struct.
func NewFlash() *FlashData {
	return (*FlashData)(web.NewFlash())
}

// Set message to flash
func (fd *FlashData) Set(key string, msg string, args ...interface{}) {
	(*web.FlashData)(fd).Set(key, msg, args...)
}

// Success writes success message to flash.
func (fd *FlashData) Success(msg string, args ...interface{}) {
	(*web.FlashData)(fd).Success(msg, args...)
}

// Notice writes notice message to flash.
func (fd *FlashData) Notice(msg string, args ...interface{}) {
	(*web.FlashData)(fd).Notice(msg, args...)
}

// Warning writes warning message to flash.
func (fd *FlashData) Warning(msg string, args ...interface{}) {
	(*web.FlashData)(fd).Warning(msg, args...)
}

// Error writes error message to flash.
func (fd *FlashData) Error(msg string, args ...interface{}) {
	(*web.FlashData)(fd).Error(msg, args...)
}

// Store does the saving operation of flash data.
// the data are encoded and saved in cookie.
func (fd *FlashData) Store(c *Controller) {
	(*web.FlashData)(fd).Store((*web.Controller)(c))
}

// ReadFromRequest parsed flash data from encoded values in cookie.
func ReadFromRequest(c *Controller) *FlashData {
	return (*FlashData)(web.ReadFromRequest((*web.Controller)(c)))
}
