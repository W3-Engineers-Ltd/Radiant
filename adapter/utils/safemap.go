package utils

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

// RadicalMap is a map with lock
type RadicalMap utils.RadicalMap

// NewRadicalMap return new safemap
func NewRadicalMap() *RadicalMap {
	return (*RadicalMap)(utils.NewRadicalMap())
}

// Get from maps return the k's value
func (m *RadicalMap) Get(k interface{}) interface{} {
	return (*utils.RadicalMap)(m).Get(k)
}

// Set Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *RadicalMap) Set(k interface{}, v interface{}) bool {
	return (*utils.RadicalMap)(m).Set(k, v)
}

// Check Returns true if k is exist in the map.
func (m *RadicalMap) Check(k interface{}) bool {
	return (*utils.RadicalMap)(m).Check(k)
}

// Delete the given key and value.
func (m *RadicalMap) Delete(k interface{}) {
	(*utils.RadicalMap)(m).Delete(k)
}

// Items returns all items in safemap.
func (m *RadicalMap) Items() map[interface{}]interface{} {
	return (*utils.RadicalMap)(m).Items()
}

// Count returns the number of items within the map.
func (m *RadicalMap) Count() int {
	return (*utils.RadicalMap)(m).Count()
}
