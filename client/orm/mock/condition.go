// Copyright 2020 radiant
//

package mock

import (
	"context"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

type Mock struct {
	cond Condition
	resp []interface{}
	cb   func(inv *orm.Invocation)
}

func NewMock(cond Condition, resp []interface{}, cb func(inv *orm.Invocation)) *Mock {
	return &Mock{
		cond: cond,
		resp: resp,
		cb:   cb,
	}
}

type Condition interface {
	Match(ctx context.Context, inv *orm.Invocation) bool
}

type SimpleCondition struct {
	tableName string
	method    string
}

func NewSimpleCondition(tableName string, methodName string) Condition {
	return &SimpleCondition{
		tableName: tableName,
		method:    methodName,
	}
}

func (s *SimpleCondition) Match(ctx context.Context, inv *orm.Invocation) bool {
	res := true
	if len(s.tableName) != 0 {
		res = res && (s.tableName == inv.GetTableName())
	}

	if len(s.method) != 0 {
		res = res && (s.method == inv.Method)
	}
	return res
}
