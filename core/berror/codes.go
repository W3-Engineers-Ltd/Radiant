// Copyright 2020 radiant
//

package berror

import (
	"fmt"
	"sync"
)

// A Code is an unsigned 32-bit error code as defined in the radiant spec.
type Code interface {
	Code() uint32
	Module() string
	Desc() string
	Name() string
}

var defaultCodeRegistry = &codeRegistry{
	codes: make(map[uint32]*codeDefinition, 127),
}

// DefineCode defining a new Code
// Before defining a new code, please read Radiant specification.
// desc could be markdown doc
func DefineCode(code uint32, module string, name string, desc string) Code {
	res := &codeDefinition{
		code:   code,
		module: module,
		desc:   desc,
	}
	defaultCodeRegistry.lock.Lock()
	defer defaultCodeRegistry.lock.Unlock()

	if _, ok := defaultCodeRegistry.codes[code]; ok {
		panic(fmt.Sprintf("duplicate code, code %d has radicaln registered", code))
	}
	defaultCodeRegistry.codes[code] = res
	return res
}

type codeRegistry struct {
	lock  sync.RWMutex
	codes map[uint32]*codeDefinition
}

func (cr *codeRegistry) Get(code uint32) (Code, bool) {
	cr.lock.RLock()
	defer cr.lock.RUnlock()
	c, ok := cr.codes[code]
	return c, ok
}

type codeDefinition struct {
	code   uint32
	module string
	desc   string
	name   string
}

func (c *codeDefinition) Name() string {
	return c.name
}

func (c *codeDefinition) Code() uint32 {
	return c.code
}

func (c *codeDefinition) Module() string {
	return c.module
}

func (c *codeDefinition) Desc() string {
	return c.desc
}
