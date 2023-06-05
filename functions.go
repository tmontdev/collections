package iterable

type PredicateHandler[T any] func(T) bool

type MapHandler[T any] func(T) any
