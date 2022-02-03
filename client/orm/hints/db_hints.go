// Copyright 2021 radiant-dev
//

package hints

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

const (
	// query level
	KeyForceIndex = iota
	KeyUseIndex
	KeyIgnoreIndex
	KeyForUpdate
	KeyLimit
	KeyOffset
	KeyOrderBy
	KeyRelDepth
)

type Hint struct {
	key   interface{}
	value interface{}
}

var _ utils.KV = new(Hint)

// GetKey return key
func (s *Hint) GetKey() interface{} {
	return s.key
}

// GetValue return value
func (s *Hint) GetValue() interface{} {
	return s.value
}

var _ utils.KV = new(Hint)

// ForceIndex return a hint about ForceIndex
func ForceIndex(indexes ...string) *Hint {
	return NewHint(KeyForceIndex, indexes)
}

// UseIndex return a hint about UseIndex
func UseIndex(indexes ...string) *Hint {
	return NewHint(KeyUseIndex, indexes)
}

// IgnoreIndex return a hint about IgnoreIndex
func IgnoreIndex(indexes ...string) *Hint {
	return NewHint(KeyIgnoreIndex, indexes)
}

// ForUpdate return a hint about ForUpdate
func ForUpdate() *Hint {
	return NewHint(KeyForUpdate, true)
}

// DefaultRelDepth return a hint about DefaultRelDepth
func DefaultRelDepth() *Hint {
	return NewHint(KeyRelDepth, true)
}

// RelDepth return a hint about RelDepth
func RelDepth(d int) *Hint {
	return NewHint(KeyRelDepth, d)
}

// Limit return a hint about Limit
func Limit(d int64) *Hint {
	return NewHint(KeyLimit, d)
}

// Offset return a hint about Offset
func Offset(d int64) *Hint {
	return NewHint(KeyOffset, d)
}

// OrderBy return a hint about OrderBy
func OrderBy(s string) *Hint {
	return NewHint(KeyOrderBy, s)
}

// NewHint return a hint
func NewHint(key interface{}, value interface{}) *Hint {
	return &Hint{
		key:   key,
		value: value,
	}
}
