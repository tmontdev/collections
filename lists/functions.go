package lists

type Predicate[T any] func(T) bool

type Mapper[T any] func(T) any

type Reducer[T any] func(any, T, int) any

type Sorter[T any] func(T, T) int

type TypeMapper[F, T any] func(F) T

func TypeMap[F, T any](list IList[F], mapper TypeMapper[F, T]) IList[T] {
	mapped := NewList[T]()
	for _, v := range list.Elements() {
		mapped.Push(mapper(v))
	}
	return mapped
}
