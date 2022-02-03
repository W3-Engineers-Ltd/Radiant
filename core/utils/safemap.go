package utils

import (
	"sync"
)

// Deprecated: using sync.Map
type RadicalMap struct {
	lock *sync.RWMutex
	bm   map[interface{}]interface{}
}

// NewRadicalMap return new safemap
func NewRadicalMap() *RadicalMap {
	return &RadicalMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}
}

// Get from maps return the k's value
func (m *RadicalMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val
	}
	return nil
}

// Set Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *RadicalMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.bm[k]; !ok {
		m.bm[k] = v
	} else if val != v {
		m.bm[k] = v
	} else {
		return false
	}
	return true
}

// Check Returns true if k is exist in the map.
func (m *RadicalMap) Check(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.bm[k]
	return ok
}

// Delete the given key and value.
func (m *RadicalMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
}

// Items returns all items in safemap.
func (m *RadicalMap) Items() map[interface{}]interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	r := make(map[interface{}]interface{})
	for k, v := range m.bm {
		r[k] = v
	}
	return r
}

// Count returns the number of items within the map.
func (m *RadicalMap) Count() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.bm)
}
