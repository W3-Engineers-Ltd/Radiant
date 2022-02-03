// Copyright 2020
//

package logs

import (
	"time"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
)

type oldToNewAdapter struct {
	old Logger
}

func (o *oldToNewAdapter) Init(config string) error {
	return o.old.Init(config)
}

func (o *oldToNewAdapter) WriteMsg(lm *logs.LogMsg) error {
	return o.old.WriteMsg(lm.When, lm.OldStyleFormat(), lm.Level)
}

func (o *oldToNewAdapter) Destroy() {
	o.old.Destroy()
}

func (o *oldToNewAdapter) Flush() {
	o.old.Flush()
}

func (o *oldToNewAdapter) SetFormatter(f logs.LogFormatter) {
	panic("unsupported operation, you should not invoke this method")
}

type newToOldAdapter struct {
	n logs.Logger
}

func (n *newToOldAdapter) Init(config string) error {
	return n.n.Init(config)
}

func (n *newToOldAdapter) WriteMsg(when time.Time, msg string, level int) error {
	return n.n.WriteMsg(&logs.LogMsg{
		When:  when,
		Msg:   msg,
		Level: level,
	})
}

func (n *newToOldAdapter) Destroy() {
	panic("implement me")
}

func (n *newToOldAdapter) Flush() {
	panic("implement me")
}
