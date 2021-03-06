// Package logs provide a general log interface
// Usage:
//
// import "github.com/W3-Engineers-Ltd/Radiant/core/logs"
//
//	log := NewLogger(10000)
//	log.SetLogger("console", "")
//
//	> the first params stand for how many channel
//
// Use it like this:
//
//	log.Trace("trace")
//	log.Info("info")
//	log.Warn("warning")
//	log.Debug("debug")
//	log.Critical("critical")
//
//  more docs http://radiant.me/docs/module/logs.md
package logs

import (
	"log"
	"time"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

// levelLogLogger is defined to implement log.Logger
// the real log level will be LevelEmergency
const levelLoggerImpl = -1

// Name for adapter with radiant official support
const (
	AdapterConsole   = "console"
	AdapterFile      = "file"
	AdapterMultiFile = "multifile"
	AdapterMail      = "smtp"
	AdapterConn      = "conn"
	AdapterEs        = "es"
	AdapterJianLiao  = "jianliao"
	AdapterSlack     = "slack"
	AdapterAliLS     = "alils"
)

// Legacy log level constants to ensure backwards compatibility.
const (
	LevelInfo  = LevelInformational
	LevelTrace = LevelDebug
	LevelWarn  = LevelWarning
)

type newLoggerFunc func() Logger

// Logger defines the behavior of a log provider.
type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

// Register makes a log provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, log newLoggerFunc) {
	logs.Register(name, func() logs.Logger {
		return &oldToNewAdapter{
			old: log(),
		}
	})
}

// RadicalLogger is default logger in radiant application.
// it can contain several providers and log message into all providers.
type RadicalLogger logs.RadicalLogger

const defaultAsyncMsgLen = 1e3

// NewLogger returns a new RadicalLogger.
// channelLen means the number of messages in chan(used where asynchronous is true).
// if the buffering chan is full, logger adapters write to file or other way.
func NewLogger(channelLens ...int64) *RadicalLogger {
	return (*RadicalLogger)(logs.NewLogger(channelLens...))
}

// Async set the log to asynchronous and start the goroutine
func (bl *RadicalLogger) Async(msgLen ...int64) *RadicalLogger {
	(*logs.RadicalLogger)(bl).Async(msgLen...)
	return bl
}

// SetLogger provides a given logger adapter into RadicalLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
func (bl *RadicalLogger) SetLogger(adapterName string, configs ...string) error {
	return (*logs.RadicalLogger)(bl).SetLogger(adapterName, configs...)
}

// DelLogger remove a logger adapter in RadicalLogger.
func (bl *RadicalLogger) DelLogger(adapterName string) error {
	return (*logs.RadicalLogger)(bl).DelLogger(adapterName)
}

func (bl *RadicalLogger) Write(p []byte) (n int, err error) {
	return (*logs.RadicalLogger)(bl).Write(p)
}

// SetLevel Set log message level.
// If message level (such as LevelDebug) is higher than logger level (such as LevelWarning),
// log providers will not even be sent the message.
func (bl *RadicalLogger) SetLevel(l int) {
	(*logs.RadicalLogger)(bl).SetLevel(l)
}

// GetLevel Get Current log message level.
func (bl *RadicalLogger) GetLevel() int {
	return (*logs.RadicalLogger)(bl).GetLevel()
}

// SetLogFuncCallDepth set log funcCallDepth
func (bl *RadicalLogger) SetLogFuncCallDepth(d int) {
	(*logs.RadicalLogger)(bl).SetLogFuncCallDepth(d)
}

// GetLogFuncCallDepth return log funcCallDepth for wrapper
func (bl *RadicalLogger) GetLogFuncCallDepth() int {
	return (*logs.RadicalLogger)(bl).GetLogFuncCallDepth()
}

// EnableFuncCallDepth enable log funcCallDepth
func (bl *RadicalLogger) EnableFuncCallDepth(b bool) {
	(*logs.RadicalLogger)(bl).EnableFuncCallDepth(b)
}

// set prefix
func (bl *RadicalLogger) SetPrefix(s string) {
	(*logs.RadicalLogger)(bl).SetPrefix(s)
}

// Emergency Log EMERGENCY level message.
func (bl *RadicalLogger) Emergency(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Emergency(format, v...)
}

// Alert Log ALERT level message.
func (bl *RadicalLogger) Alert(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Alert(format, v...)
}

// Critical Log CRITICAL level message.
func (bl *RadicalLogger) Critical(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Critical(format, v...)
}

// Error Log ERROR level message.
func (bl *RadicalLogger) Error(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Error(format, v...)
}

// Warning Log WARNING level message.
func (bl *RadicalLogger) Warning(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Warning(format, v...)
}

// Notice Log NOTICE level message.
func (bl *RadicalLogger) Notice(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Notice(format, v...)
}

// Informational Log INFORMATIONAL level message.
func (bl *RadicalLogger) Informational(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Informational(format, v...)
}

// Debug Log DEBUG level message.
func (bl *RadicalLogger) Debug(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Debug(format, v...)
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func (bl *RadicalLogger) Warn(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Warn(format, v...)
}

// Info Log INFO level message.
// compatibility alias for Informational()
func (bl *RadicalLogger) Info(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Info(format, v...)
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func (bl *RadicalLogger) Trace(format string, v ...interface{}) {
	(*logs.RadicalLogger)(bl).Trace(format, v...)
}

// Flush flush all chan data.
func (bl *RadicalLogger) Flush() {
	(*logs.RadicalLogger)(bl).Flush()
}

// Close close logger, flush all chan data and destroy all adapters in RadicalLogger.
func (bl *RadicalLogger) Close() {
	(*logs.RadicalLogger)(bl).Close()
}

// Reset close all outputs, and set bl.outputs to nil
func (bl *RadicalLogger) Reset() {
	(*logs.RadicalLogger)(bl).Reset()
}

// GetRadicalLogger returns the default RadicalLogger
func GetRadicalLogger() *RadicalLogger {
	return (*RadicalLogger)(logs.GetRadicalLogger())
}

// GetLogger returns the default RadicalLogger
func GetLogger(prefixes ...string) *log.Logger {
	return logs.GetLogger(prefixes...)
}

// Reset will remove all the adapter
func Reset() {
	logs.Reset()
}

// Async set the radicallogger with Async mode and hold msglen messages
func Async(msgLen ...int64) *RadicalLogger {
	return (*RadicalLogger)(logs.Async(msgLen...))
}

// SetLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	logs.SetLevel(l)
}

// SetPrefix sets the prefix
func SetPrefix(s string) {
	logs.SetPrefix(s)
}

// EnableFuncCallDepth enable log funcCallDepth
func EnableFuncCallDepth(b bool) {
	logs.EnableFuncCallDepth(b)
}

// SetLogFuncCall set the CallDepth, default is 4
func SetLogFuncCall(b bool) {
	logs.SetLogFuncCall(b)
}

// SetLogFuncCallDepth set log funcCallDepth
func SetLogFuncCallDepth(d int) {
	logs.SetLogFuncCallDepth(d)
}

// SetLogger sets a new logger.
func SetLogger(adapter string, config ...string) error {
	return logs.SetLogger(adapter, config...)
}

// Emergency logs a message at emergency level.
func Emergency(f interface{}, v ...interface{}) {
	logs.Emergency(f, v...)
}

// Alert logs a message at alert level.
func Alert(f interface{}, v ...interface{}) {
	logs.Alert(f, v...)
}

// Critical logs a message at critical level.
func Critical(f interface{}, v ...interface{}) {
	logs.Critical(f, v...)
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	logs.Error(f, v...)
}

// Warning logs a message at warning level.
func Warning(f interface{}, v ...interface{}) {
	logs.Warning(f, v...)
}

// Warn compatibility alias for Warning()
func Warn(f interface{}, v ...interface{}) {
	logs.Warn(f, v...)
}

// Notice logs a message at notice level.
func Notice(f interface{}, v ...interface{}) {
	logs.Notice(f, v...)
}

// Informational logs a message at info level.
func Informational(f interface{}, v ...interface{}) {
	logs.Informational(f, v...)
}

// Info compatibility alias for Warning()
func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v...)
}

// Debug logs a message at debug level.
func Debug(f interface{}, v ...interface{}) {
	logs.Debug(f, v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Trace(f interface{}, v ...interface{}) {
	logs.Trace(f, v...)
}

func init() {
	SetLogFuncCallDepth(4)
}
