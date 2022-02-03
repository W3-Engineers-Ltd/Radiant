// Copyright 2020 radiant
//

package orm

import (
	"context"
	"time"
)

// Invocation represents an "Orm" invocation
type Invocation struct {
	Method string
	// Md may be nil in some cases. It depends on method
	Md interface{}
	// the args are all arguments except context.Context
	Args []interface{}

	mi *modelInfo
	// f is the Orm operation
	f func(ctx context.Context) []interface{}

	// insideTx indicates whether this is inside a transaction
	InsideTx    bool
	TxStartTime time.Time
	TxName      string
}

func (inv *Invocation) GetTableName() string {
	if inv.mi != nil {
		return inv.mi.table
	}
	return ""
}

func (inv *Invocation) execute(ctx context.Context) []interface{} {
	return inv.f(ctx)
}

// GetPkFieldName return the primary key of this table
// if not found, "" is returned
func (inv *Invocation) GetPkFieldName() string {
	if inv.mi.fields.pk != nil {
		return inv.mi.fields.pk.name
	}
	return ""
}
