package collections

// MapFrom returns a new HashMap from the given built-in source map.
// Changes in the returned HashMap will not affect the source map
func MapFrom[K comparable, V any](source map[K]V) Map[K, V] {
	return HashMap[K, V](source).Clone()
}

// HashMap is the default implementation of Map. Consider it a built-in map (but super powered)
type HashMap[K comparable, V any] map[K]V

// Length returns how many values are stored in the HashMap.
func (m HashMap[K, V]) Length() int {
	return len(m)
}

// IsEmpty returns true if there are *no* value stored in the HashMap.
func (m HashMap[K, V]) IsEmpty() bool {
	return m.Length() == 0
}

// IsNotEmpty returns true if there are values stored in the HashMap.
func (m HashMap[K, V]) IsNotEmpty() bool {
	return !m.IsEmpty()
}

// Where returns a new Map containing only the key/value which satisfies de Predicate
func (m HashMap[K, V]) Where(predicate KeyValuePredicate[K, V]) Map[K, V] {
	filtered := HashMap[K, V]{}
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
func (m HashMap[K, V]) RemoveWhere(predicate KeyValuePredicate[K, V]) Map[K, V] {
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

// Some returns true if one or more key/value stored in HashMap satisfies the Predicate
func (m HashMap[K, V]) Some(predicate KeyValuePredicate[K, V]) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// None returns true if *no* key/value stored in the HashMap satisfies the Predicate.
func (m HashMap[K, V]) None(predicate KeyValuePredicate[K, V]) bool {
	return !m.Some(predicate)
}

// Every returns true if every value stored in the HashMap satisfies the predicate.
func (m HashMap[K, V]) Every(predicate KeyValuePredicate[K, V]) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// Set sets the given value in the given key, and then returns itself
func (m HashMap[K, V]) Set(key K, value V) Map[K, V] {
	m[key] = value
	return m
}

// Get returns the value stored in the given key from the HashMap
func (m HashMap[K, V]) Get(key K) V {
	return m[key]
}

// Access returns the value stored in the given key (if stored)
func (m HashMap[K, V]) Access(key K) (V, bool) {
	value, has := m[key]
	return value, has
}

// Clone returns a new HashMap with the same keys and values from the original
func (m HashMap[K, V]) Clone() Map[K, V] {
	cloned := HashMap[K, V]{}
	for k, v := range m {
		cloned.Set(k, v)
	}
	return cloned
}

// Has returns true if the given key is filled.
func (m HashMap[K, V]) Has(key K) bool {
	_, has := m[key]
	return has
}

// Keys returns a SimpleList with all keys
func (m HashMap[K, V]) Keys() List[K] {
	list := &SimpleList[K]{}
	for k := range m {
		list.Push(k)
	}
	return list
}

// Values returns a SimpleList with all values
func (m HashMap[K, V]) Values() List[V] {
	list := &SimpleList[V]{}
	for _, v := range m {
		list.Push(v)
	}
	return list
}

// Merge sets all key/value pairs from the given map in itself.
// Ignores key conflicts when replace is false
func (m HashMap[K, V]) Merge(from Map[K, V], replace bool) Map[K, V] {
	for k, v := range from.Builtin() {
		if !m.Has(k) || replace {
			m.Set(k, v)
		}
	}
	return m
}

func (m HashMap[K, V]) Builtin() map[K]V {
	return m
}

func (m HashMap[K, V]) HashMap() HashMap[K, V] {
	return m
}
