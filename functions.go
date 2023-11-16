package collections

type Predicate[T any] func(T) bool

type KeyValuePredicate[K comparable, V any] func(K, V) bool

type Mapper[T any] func(T) any

type Reducer[T any] func(any, T, int) any

type Sorter[T any] func(T, T) int
