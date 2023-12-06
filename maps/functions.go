package maps

type Predicate[K comparable, V any] func(K, V) bool
