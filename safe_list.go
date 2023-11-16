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

func (s *SafeList[T]) self(exec func() any) *SafeList[T] {
	protect[any, T](s, exec)
	return s
}

// NewSafeList returns a new SafeList with the given elements
func NewSafeList[T any](elements ...T) *SafeList[T] {
	return &SafeList[T]{l: NewDynamicList(elements...)}
}

// NewSafeListFrom returns a new SafeList with the given slice
func NewSafeListFrom[T any](elements []T) *SafeList[T] {
	return &SafeList[T]{l: NewDynamicListFrom(elements)}
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
	return protect[T, T](s, func() T {
		return s.l.ElementAt(i)
	})
}

// Elements returns a built-in slice with all elements in the SafeList.
func (s *SafeList[T]) Elements() []T {
	return protect[any, T](s, func() any {
		return s.l.Elements()
	}).([]T)
}

// Push add the given elements in the SafeList, and then returns itself.
func (s *SafeList[T]) Push(elements ...T) List[T] {
	return s.self(func() any {
		return s.l.Push(elements...)
	})
}

// Clone returns an identical SafeList from the original.
func (s *SafeList[T]) Clone() List[T] {
	return &SafeList[T]{l: protect[List[T], T](s, func() List[T] {
		return s.l.Clone()
	}).(*DynamicList[T])}
}

// FirstElement returns the first element in the SafeList.
// If SafeList is empty (see IsEmpty), panics
func (s *SafeList[T]) FirstElement() T {
	return protect[T, T](s, func() T {
		return s.l.FirstElement()
	})
}

// First returns the pointer of the first element in the SafeList.
// If SafeList is empty (see IsEmpty), nil will be returned.
func (s *SafeList[T]) First() *T {
	return protect[*T, T](s, func() *T {
		return s.l.First()
	})
}

// LastElement returns the last element in the SafeList.
// If SafeList is empty (see IsEmpty), panics.
func (s *SafeList[T]) LastElement() T {
	return protect[T, T](s, func() T {
		return s.l.LastElement()
	})
}

// Last returns the pointer of the last element in the SafeList.
// If SafeList is empty (see IsEmpty), nil will be returned.
func (s *SafeList[T]) Last() *T {
	return protect[*T, T](s, func() *T {
		return s.l.Last()
	})
}

// FirstIndexWhere returns the index of the first element which satisfies the predicate.
// If no element satisfies the predicate, -1 will be returned.
func (s *SafeList[T]) FirstIndexWhere(handler Predicate[T]) int {
	return protect[int, T](s, func() int {
		return s.l.FirstIndexWhere(handler)
	})
}

// FirstWhere returns the pointer of the first element which satisfies the predicate.
// If no element satisfies the predicate, nil will be returned.
func (s *SafeList[T]) FirstWhere(handler Predicate[T]) *T {
	return protect[*T, T](s, func() *T {
		return s.l.FirstWhere(handler)
	})
}

// FirstElementWhere returns the first element which satisfies the predicate.
// If no element satisfies the predicate, panics.
func (s *SafeList[T]) FirstElementWhere(handler Predicate[T]) T {
	return protect[T, T](s, func() T {
		return s.l.FirstElementWhere(handler)
	})
}

// LastIndexWhere returns the index of the last element which satisfies the predicate.
// If no element satisfies the predicate, -1 will be returned.
func (s *SafeList[T]) LastIndexWhere(handler Predicate[T]) int {
	return protect[int, T](s, func() int {
		return s.l.LastIndexWhere(handler)
	})
}

// LastWhere returns the pointer of the last element which satisfies the predicate.
// If no element satisfies the predicate, nil will be returned.
func (s *SafeList[T]) LastWhere(handler Predicate[T]) *T {
	return protect[*T, T](s, func() *T {
		return s.l.LastWhere(handler)
	})
}

// LastElementWhere returns the last element which satisfies the predicate.
// If no element satisfies the predicate, panics.
func (s *SafeList[T]) LastElementWhere(handler Predicate[T]) T {
	return protect[T, T](s, func() T {
		return s.l.LastElementWhere(handler)
	})
}

// IndexWhere returns a DynamicList[int] for all element index which satisfies the predicate.
// If no element satisfies the predicate, an empty DynamicList will be returned.
func (s *SafeList[T]) IndexWhere(handler Predicate[T]) List[int] {
	return protect[List[int], T](s, func() List[int] {
		return s.l.IndexWhere(handler)
	})
}

// Where returns a DynamicList with all the elements which satisfies the predicate.
// If no element satisfies the predicate, an empty DynamicList will be returned.
func (s *SafeList[T]) Where(handler Predicate[T]) List[T] {
	return protect[List[T], T](s, func() List[T] {
		return s.l.Where(handler)
	})
}

// Map iterates over the element of the SafeList calling Mapper, and return a new DynamicList with the results.
func (s *SafeList[T]) Map(handler Mapper[T]) List[any] {
	return protect[List[any], T](s, func() List[any] {
		return s.l.Map(handler)
	})
}

// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be the accumulator for the next.
// The final result will be returned.
func (s *SafeList[T]) Reduce(reducer Reducer[T], accumulator any) any {
	return protect[any, T](s, func() any {
		return s.l.Reduce(reducer, accumulator)
	})
}

// Every returns true if every element in the List satisfies the predicate.
func (s *SafeList[T]) Every(handler Predicate[T]) bool {
	return protect[bool, T](s, func() bool {
		return s.l.Every(handler)
	})
}

// Some returns true if at least one element in the List satisfies the predicate.
func (s *SafeList[T]) Some(handler Predicate[T]) bool {
	return protect[bool, T](s, func() bool {
		return s.l.Some(handler)
	})
}

// None returns true no element in the List satisfy the predicate.
func (s *SafeList[T]) None(handler Predicate[T]) bool {
	return protect[bool, T](s, func() bool {
		return s.l.None(handler)
	})
}

// Pop removes the last element from the List and returns itself.
func (s *SafeList[T]) Pop() List[T] {
	return s.self(func() any {
		return s.l.Pop()
	})
}

// Shift removes the first element from the List and then returns itself.
func (s *SafeList[T]) Shift() List[T] {
	return s.self(func() any {
		return s.l.Shift()
	})
}

// Set sets the given element at the given index, and then returns itself.
func (s *SafeList[T]) Set(index int, element T) List[T] {
	return s.self(func() any {
		return s.l.Set(index, element)
	})
}

// Interval returns a new DynamicList with all elements between the *from* and *to* indexes.
func (s *SafeList[T]) Interval(from, to int) List[T] {
	return protect[List[T], T](s, func() List[T] {
		return s.l.Interval(from, to)
	})
}

// String returns a string representation of the SafeList.
func (s *SafeList[T]) String() string {
	return protect[string, T](s, func() string {
		return s.l.String()
	})
}

// Sort receives a Sorter function to sort its elements, and returns itself after sorted.
func (s *SafeList[T]) Sort(sorter Sorter[T]) List[T] {
	return s.self(func() any {
		return s.l.Sort(sorter)
	})
}

// Clear removes all elements from the SafeList, making it empty, and then returns itself.
func (s *SafeList[T]) Clear() List[T] {
	return s.self(func() any {
		return s.l.Clear()
	})
}

// IsDynamicallySized returns true, as SafeList is a dynamically-sized implementation of List
func (s *SafeList[T]) IsDynamicallySized() bool {
	return true
}

// IsThreadSafe returns true, as SafeList is a thread-safe implementation of List
func (s *SafeList[T]) IsThreadSafe() bool {
	return true
}
