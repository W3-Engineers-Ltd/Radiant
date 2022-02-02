package orm

import (
	"io"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

// Log implement the log.Logger
type Log orm.Log

// costomer log func
var LogFunc = orm.LogFunc

// NewLog set io.Writer to create a Logger.
func NewLog(out io.Writer) *Log {
	return (*Log)(orm.NewLog(out))
}
