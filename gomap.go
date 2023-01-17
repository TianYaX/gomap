package gomap

import (
	"sync"
)

type SMap[K comparable, V any] struct {
	*sync.Map
}

func Make[K comparable, V any]() SMap[K, V] {
	return SMap[K, V]{new(sync.Map)}
}

func (m *SMap[k, v]) Init() {
	if m.Map == nil {
		m.Map = new(sync.Map)
	}
}

func (m SMap[K, V]) Get(key K) V {
	val, ok := m.Load(key)
	if !ok {
		var val V
		return val
	}
	return val.(V)
}

func (m SMap[K, V]) Put(key K, val V) {
	m.Store(key, val)
}

func (m SMap[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

func (m SMap[K, V]) Range(f func(key K, val V) bool) {
	m.Map.Range(func(key, val any) bool {
		return f(key.(K), val.(V))
	})
}

func (m SMap[K, V]) TryPut(key K, val V) (V, bool) {
	ret, ok := m.LoadOrStore(key, val)
	return ret.(V), ok
}

func (m SMap[K, V]) TryDelete(key K) (V, bool) {
	val, ok := m.LoadAndDelete(key)
	return val.(V), ok
}

type Map[K comparable, V any] struct {
	mp map[K]V
	mx *sync.RWMutex
}

func MakeMap[K comparable, V any](num int) Map[K, V] {
	return Map[K, V]{make(map[K]V, num), new(sync.RWMutex)}
}

func (m Map[K, V]) Get(key K) V {
	m.mx.RLock()
	defer m.mx.RUnlock()
	return m.mp[key]
}

func (m Map[K, V]) Put(key K, val V) {
	m.mx.Lock()
	m.mp[key] = val
	m.mx.Unlock()
}

func (m Map[K, V]) Range(f func(key K, val V) bool) {
	for k, v := range m.mp {
		if !f(k, v) {
			break
		}
	}
}
