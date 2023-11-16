package collections

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) IsEmpty() bool {
	return m.Length() == 0
}

func (m Map[K, V]) IsNotEmpty() bool {
	return !m.IsEmpty()
}

func (m Map[K, V]) Length() int {
	return len(m)
}

func (m Map[K, V]) Where(predicate KeyValuePredicate[K, V]) Dictionary[K, V] {
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

func (m Map[K, V]) RemoveWhere(predicate KeyValuePredicate[K, V]) Dictionary[K, V] {
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

func (m Map[K, V]) Some(predicate KeyValuePredicate[K, V]) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

func (m Map[K, V]) None(predicate KeyValuePredicate[K, V]) bool {
	return !m.Some(predicate)
}

func (m Map[K, V]) Every(predicate KeyValuePredicate[K, V]) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

func (m Map[K, V]) Set(key K, value V) Dictionary[K, V] {
	m[key] = value
	return m
}

func (m Map[K, V]) Get(key K) V {
	return m[key]
}

func (m Map[K, V]) Access(key K) (V, bool) {
	value, has := m[key]
	return value, has
}

func (m Map[K, V]) Has(key K) bool {
	_, has := m[key]
	return has
}

func (m Map[K, V]) Clone() Dictionary[K, V] {
	cloned := Map[K, V]{}
	for k, v := range m {
		cloned.Set(k, v)
	}
	return cloned
}

func (m Map[K, V]) Keys() List[K] {
	list := &SimpleList[K]{}
	for k, _ := range m {
		list.Push(k)
	}
	return list
}

func (m Map[K, V]) Values() List[V] {
	list := &SimpleList[V]{}
	for _, v := range m {
		list.Push(v)
	}
	return list
}

func (m Map[K, V]) Merge(from Map[K, V], replace bool) Map[K, V] {
	for k, v := range from {
		if !m.Has(k) || replace {
			m.Set(k, v)
		}
	}
	return m
}
