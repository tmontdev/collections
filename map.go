package collection

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

func (m Map[K, V]) Where(predicate KeyPredicate[K, V]) Dictionary[K, V] {
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

func (m Map[K, V]) RemoveWhere(predicate KeyPredicate[K, V]) Dictionary[K, V] {
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

func (m Map[K, V]) Some(predicate KeyPredicate[K, V]) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

func (m Map[K, V]) None(predicate KeyPredicate[K, V]) bool {
	return !m.Some(predicate)
}

func (m Map[K, V]) Every(predicate KeyPredicate[K, V]) bool {
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
