package collections

import (
	"fmt"
)

// NewList returns a new List with the given elements
func NewList[T any](elements ...T) *List[T] {
	l := &List[T]{}
	return l.Push(elements...).(*List[T])
}

type List[T any] []T

// Length method returns how many element are in the List
func (l *List[T]) Length() int {
	return len(l.Elements())
}

// IsEmpty returns true if Length is zero
func (l *List[T]) IsEmpty() bool {
	return l.Length() == 0
}

// IsNotEmpty returns true if Length is not zero
func (l *List[T]) IsNotEmpty() bool {
	return !l.IsEmpty()
}

// At returns the pointer of the element in the given index from the List
// If there is no element in the index, nil will be returned
func (l *List[T]) At(i int) (at *T) {
	if l.Length()-1 < i || i < 0 {
		return
	}
	at = &l.Elements()[i]
	return
}

// ElementAt returns the element in the given index from the List
// If there is no element in the index, panics
func (l *List[T]) ElementAt(i int) T {
	return l.Elements()[i]
}

// Elements returns a built-in slice with all elements in the List
func (l *List[T]) Elements() []T {
	return *l
}

// Push add the given elements in the List, and returns itself
func (l *List[T]) Push(elements ...T) Iterable[T] {
	*l = append(l.Elements(), elements...)
	return l
}

// Clone returns an identical List from the original
func (l *List[T]) Clone() Iterable[T] {
	return NewList[T](l.Elements()...)
}

// FirstElement returns the first element in the List
// if isEmpty, panics
func (l *List[T]) FirstElement() T {
	return l.ElementAt(0)
}

// First returns the pointer of the first element in the List
// if isEmpty, nil will be returned
func (l *List[T]) First() *T {
	return l.At(0)
}

// LastElement returns the last element in the List
// if isEmpty, panics
func (l *List[T]) LastElement() T {
	return l.ElementAt(l.Length() - 1)
}

// Last returns the pointer of the last element in the List
// if isEmpty, nil will be returned
func (l *List[T]) Last() *T {
	return l.At(l.Length() - 1)
}

// FirstIndexWhere returns the index of the first element witch satisfies the predicate
// if no element satisfies the predicate, -1 will be returned
func (l *List[T]) FirstIndexWhere(handler Predicate[T]) int {
	for i, v := range l.Elements() {
		if handler(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere returns the index of the last element witch satisfies the predicate
// if no element satisfies the predicate, -1 will be returned
func (l *List[T]) LastIndexWhere(handler Predicate[T]) int {
	idx := -1
	for i, v := range l.Elements() {
		if handler(v) {
			idx = i
		}
	}
	return idx
}

// IndexWhere returns a new List[int] for all element index witch satisfies the predicate
// if no element satisfies the predicate, an empty List will be returned
func (l *List[T]) IndexWhere(handler Predicate[T]) Iterable[int] {
	list := NewList[int]()
	for i, v := range l.Elements() {
		if handler(v) {
			list.Push(i)
		}
	}
	return list
}

// Where returns a new List with all the elements witch satisfies the predicate.
// if no element satisfies the predicate, an empty List will be returned
func (l *List[T]) Where(handler Predicate[T]) Iterable[T] {
	selected := NewList[T]()
	for _, v := range l.Elements() {
		if handler(v) {
			selected.Push(v)
		}
	}
	return selected
}

// Map iterates over the elements of the List calling Mapper, and return a new List with the results.
func (l *List[T]) Map(handler Mapper[T]) Iterable[any] {
	mapped := NewList[any]()
	for _, v := range l.Elements() {
		mapped.Push(handler(v))
	}
	return mapped
}

// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be accumulator for the next
// The final result will be returned
func (l *List[T]) Reduce(reducer Reducer[T], accumulator any) any {
	for i, v := range l.Elements() {
		accumulator = reducer(accumulator, v, i)
	}
	return accumulator
}

// Every returns true if every element in the List satisfy the predicate
func (l *List[T]) Every(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if !handler(v) {
			return false
		}
	}
	return true
}

// Some returns true if at least one element in the List satisfy the predicate
func (l *List[T]) Some(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if handler(v) {
			return true
		}
	}
	return false
}

// None returns true no element in the List satisfy the predicate
func (l *List[T]) None(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if handler(v) {
			return false
		}
	}
	return true
}

// Pop removes the last element from the List and returns itself
func (l *List[T]) Pop() Iterable[T] {
	*l = l.Elements()[0 : l.Length()-1]
	return l
}

// Shift removes the first element from the List and returns itself
func (l *List[T]) Shift() Iterable[T] {
	*l = l.Elements()[1:l.Length()]
	return l
}

// Set sets the given element in the given index, and returns itself
// if the given index is not yet filled, panics
func (l *List[T]) Set(index int, element T) Iterable[T] {
	l.ElementAt(index)
	at := l.At(index)
	*at = element
	return l
}

// Interval returns a new Iterable with all elements between from and to given indexes
func (l *List[T]) Interval(from, to int) Iterable[T] {
	return NewList[T](l.Elements()[from : to+1]...)
}

// String returns a string representation of the List
func (l *List[T]) String() string {
	return fmt.Sprint(l.Elements())
}

// Sort receive a Sorter function to sort its elements, and returns itself after sorted
func (l *List[T]) Sort(sorter Sorter[T]) Iterable[T] {
	changed := true
	for changed {
		changed = l.sort(sorter)
	}
	return l
}

func (l *List[T]) sort(sorter Sorter[T]) bool {
	changed := false
	for i, _ := range l.Elements() {
		if i == 0 {
			continue
		}
		a, b := l.ElementAt(i-1), l.ElementAt(i)
		step := sorter(a, b)
		if step <= 0 {
			continue
		}
		if step > 0 {
			l.Set(i, a)
			l.Set(i-1, b)
			changed = true
		}
	}
	return changed
}
