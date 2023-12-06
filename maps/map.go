package maps

import (
	"encoding/json"
	"github.com/tmontdev/collections/lists"
)

// From returns a new Map from the given built-in source map.
// Changes in the returned Map will not affect the source map
func From[K comparable, V any](source map[K]V) IMap[K, V] {
	return Map[K, V](source).Clone()
}

// Map is the default implementation of IMap. Consider it a built-in map (but super powered)
type Map[K comparable, V any] map[K]V

// Length returns how many values are stored in the Map.
func (m Map[K, V]) Length() int {
	return len(m)
}

// IsEmpty returns true if there are *no* value stored in the Map.
func (m Map[K, V]) IsEmpty() bool {
	return m.Length() == 0
}

// IsNotEmpty returns true if there are values stored in the Map.
func (m Map[K, V]) IsNotEmpty() bool {
	return !m.IsEmpty()
}

// Where returns a new IMap containing only the key/value which satisfies de Predicate
func (m Map[K, V]) Where(predicate Predicate[K, V]) IMap[K, V] {
	filtered := Map[K, V]{}
	if m.IsEmpty() {
		return filtered
	}
	for k, v := range m {
		if predicate(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

// RemoveWhere deletes all key/value which satisfies the Predicate, and then returns itself
func (m Map[K, V]) RemoveWhere(predicate Predicate[K, V]) IMap[K, V] {
	if m.IsEmpty() {
		return m
	}
	for k, v := range m {
		if predicate(k, v) {
			delete(m, k)
		}
	}
	return m
}

// Some returns true if one or more key/value stored in Map satisfies the Predicate
func (m Map[K, V]) Some(predicate Predicate[K, V]) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// None returns true if *no* key/value stored in the Map satisfies the Predicate.
func (m Map[K, V]) None(predicate Predicate[K, V]) bool {
	return !m.Some(predicate)
}

// Every returns true if every value stored in the Map satisfies the predicate.
func (m Map[K, V]) Every(predicate Predicate[K, V]) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// Set sets the given value in the given key, and then returns itself
func (m Map[K, V]) Set(key K, value V) IMap[K, V] {
	m[key] = value
	return m
}

// Get returns the value stored in the given key from the Map
func (m Map[K, V]) Get(key K) V {
	return m[key]
}

// Access returns the value stored in the given key (if stored)
func (m Map[K, V]) Access(key K) (V, bool) {
	value, has := m[key]
	return value, has
}

// Clone returns a new Map with the same keys and values from the original
func (m Map[K, V]) Clone() IMap[K, V] {
	cloned := Map[K, V]{}
	for k, v := range m {
		cloned.Set(k, v)
	}
	return cloned
}

// Has returns true if the given key is filled.
func (m Map[K, V]) Has(key K) bool {
	_, has := m[key]
	return has
}

// Keys returns a List with all keys
func (m Map[K, V]) Keys() lists.IList[K] {
	list := &lists.List[K]{}
	for k := range m {
		list.Push(k)
	}
	return list
}

// Values returns a List with all values
func (m Map[K, V]) Values() lists.IList[V] {
	list := &lists.List[V]{}
	for _, v := range m {
		list.Push(v)
	}
	return list
}

// Complement sets missing key/value pairs from the given map in itself.
func (m Map[K, V]) Complement(from IMap[K, V]) IMap[K, V] {
	for k, v := range from.Builtin() {
		if !m.Has(k) {
			m.Set(k, v)
		}
	}
	return m
}

// SetFrom sets all key/value pairs from the given map in itself.
func (m Map[K, V]) SetFrom(from IMap[K, V]) IMap[K, V] {
	for k, v := range from.Builtin() {
		m.Set(k, v)
	}
	return m
}

func (m Map[K, V]) Builtin() map[K]V {
	return m
}

func (m Map[K, V]) HashMap() Map[K, V] {
	return m
}

func (m Map[K, V]) Struct(str any) error {
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, str)
	return err
}
