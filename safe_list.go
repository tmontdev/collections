package collections

import "sync"

type SafeList[T any] struct {
	l *DynamicList[T]
	sync.Mutex
}

func protect[T, U any](list *SafeList[U], exec func() T) T {
	list.Lock()
	defer list.Unlock()
	return exec()
}

func NewSafeList[T any](elements ...T) *SafeList[T] {
	return &SafeList[T]{l: NewDynamicList(elements...)}
}

func NewSafeListFrom[T any](elements []T) *SafeList[T] {
	return &SafeList[T]{l: NewDynamicListFrom(elements)}
}

//func (s *SafeList[T]) protect(exec func() any) any {
//	s.Lock()
//	defer s.Unlock()
//	return exec()
//}

func (s *SafeList[T]) self(exec func() any) *SafeList[T] {
	protect[any, T](s, exec)
	return s
}

// Length returns how many elements are in the SafeList.
func (s *SafeList[T]) Length() int {
	return protect[int, T](s, func() int {
		return s.l.Length()
	})
}

// IsEmpty returns true if there are *no* Elements stored in the SafeList.
func (s *SafeList[T]) IsEmpty() bool {
	return protect[bool, T](s, func() bool {
		return s.l.IsEmpty()
	})
}

// IsNotEmpty returns true if there are Elements stored in the SafeList.
func (s *SafeList[T]) IsNotEmpty() bool {
	return protect[bool, T](s, func() bool {
		return s.l.IsNotEmpty()
	})
}

// At returns the pointer of the element at the given index from the SafeList.
// If there is no element at the given index, nil will be returned.
func (s *SafeList[T]) At(i int) *T {
	return protect[any, T](s, func() any {
		return s.l.At(i)
	}).(*T)
}

// ElementAt returns the element at the given index from the SafeList.
// If there is no element at the given index, panics.
func (s *SafeList[T]) ElementAt(i int) T {
	return protect[any, T](s, func() any {
		return s.l.ElementAt(i)
	})
}

// Elements returns a built-in slice with all elements in the SafeList.
func (s *SafeList[T]) Elements() []T {
	return protect[any, T](s, func() any {
		return s.l.Elements()
	}).([]T)
}

// Push add the given elements in the DynamicList, and then returns itself.
func (s *SafeList[T]) Push(elements ...T) List[T] {
	return s.self(func() any {
		return s.l.Push(elements...)
	})
}

func (s *SafeList[T]) Clone() List[T] {
	return &SafeList[T]{l: protect[List[T], T](s, func() List[T] {
		return s.l.Clone()
	}).(*DynamicList[T])}
}

func (s *SafeList[T]) FirstElement() T {
	return protect[any, T](s, func() any {
		return s.l.FirstElement()
	})
}

func (s *SafeList[T]) First() *T {
	return protect[*T, T](s, func() *T {
		return s.l.First()
	})
}

func (s *SafeList[T]) LastElement() T {
	return protect[T, T](s, func() T {
		return s.l.LastElement()
	})
}

func (s *SafeList[T]) Last() *T {
	return protect[*T, T](s, func() *T {
		return s.l.Last()
	})
}

func (s *SafeList[T]) FirstIndexWhere(handler Predicate[T]) int {
	return protect[int, T](s, func() int {
		return s.l.FirstIndexWhere(handler)
	})
}

func (s *SafeList[T]) FirstWhere(handler Predicate[T]) *T {
	return protect[*T, T](s, func() *T {
		return s.l.FirstWhere(handler)
	})
}

func (s *SafeList[T]) FirstElementWhere(handler Predicate[T]) T {
	return protect[T, T](s, func() T {
		return s.l.FirstElementWhere(handler)
	})
}

func (s *SafeList[T]) LastIndexWhere(handler Predicate[T]) int {
	return protect[int, T](s, func() int {
		return s.l.LastIndexWhere(handler)
	})
}

func (s *SafeList[T]) LastWhere(handler Predicate[T]) *T {
	return protect[*T, T](s, func() *T {
		return s.l.LastWhere(handler)
	})
}

func (s *SafeList[T]) LastElementWhere(handler Predicate[T]) T {
	return protect[T, T](s, func() T {
		return s.l.LastElementWhere(handler)
	})
}

func (s *SafeList[T]) IndexWhere(handler Predicate[T]) List[int] {
	return protect[List[int], T](s, func() List[int] {
		return s.l.IndexWhere(handler)
	})
}

func (s *SafeList[T]) Where(handler Predicate[T]) List[T] {
	return protect[List[T], T](s, func() List[T] {
		return s.l.Where(handler)
	})
}

func (s *SafeList[T]) Map(handler Mapper[T]) List[any] {
	return protect[List[any], T](s, func() List[any] {
		return s.l.Map(handler)
	})
}

func (s *SafeList[T]) Reduce(reducer Reducer[T], accumulator any) any {
	return protect[any, T](s, func() any {
		return s.l.Reduce(reducer, accumulator)
	})
}

func (s *SafeList[T]) Every(handler Predicate[T]) bool {
	return protect[bool, T](s, func() bool {
		return s.l.Every(handler)
	})
}

func (s *SafeList[T]) Some(handler Predicate[T]) bool {
	return protect[bool, T](s, func() bool {
		return s.l.Some(handler)
	})
}

func (s *SafeList[T]) None(handler Predicate[T]) bool {
	return protect[bool, T](s, func() bool {
		return s.l.None(handler)
	})
}

func (s *SafeList[T]) Pop() List[T] {
	return s.self(func() any {
		return s.l.Pop()
	})
}

func (s *SafeList[T]) Shift() List[T] {
	return s.self(func() any {
		return s.l.Shift()
	})
}

func (s *SafeList[T]) Set(index int, element T) List[T] {
	return s.self(func() any {
		return s.l.Set(index, element)
	})
}

func (s *SafeList[T]) Interval(from, to int) List[T] {
	return protect[List[T], T](s, func() List[T] {
		return s.l.Interval(from, to)
	})
}

func (s *SafeList[T]) String() string {
	return protect[string, T](s, func() string {
		return s.l.String()
	})
}

func (s *SafeList[T]) Sort(sorter Sorter[T]) List[T] {
	return s.self(func() any {
		return s.l.Sort(sorter)
	})
}

func (s *SafeList[T]) Clear() List[T] {
	return s.self(func() any {
		return s.l.Clear()
	})
}
