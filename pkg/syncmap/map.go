// generate map with strong types:
// # go get github.com/cheekybits/genny
// # cat map.go | genny gen "KeyType=string ValueType=uint64"
package syncmap

import "sync"
import "github.com/cheekybits/genny/generic"

type KeyType generic.Type
type ValueType generic.Type

type MapKeyTypeValueType interface {
	Load(key KeyType) (ValueType, bool)
	Store(key KeyType, val ValueType)
	Delete(key KeyType)
}

type mapKeyTypeValueType struct {
	data  map[KeyType]ValueType
	mutex sync.RWMutex
}

func NewMapKeyTypeValueType() MapKeyTypeValueType {
	return &mapKeyTypeValueType{
		data: make(map[KeyType]ValueType),
	}
}

func (m *mapKeyTypeValueType) Load(key KeyType) (ValueType, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	val, has := m.data[key]
	return val, has
}

func (m *mapKeyTypeValueType) Store(key KeyType, val ValueType) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[key] = val
}

func (m *mapKeyTypeValueType) Delete(key KeyType) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, key)
}
