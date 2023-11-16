package collections

import (
	"fmt"
)

// NewDynamicList returns a new DynamicList with the given elements
func NewDynamicList[T any](elements ...T) *DynamicList[T] {
	l := &DynamicList[T]{}
	l.Push(elements...)
	return l
}

// NewDynamicListFrom returns a new DynamicList with the given elements
func NewDynamicListFrom[T any](elements []T) *DynamicList[T] {
	l := &DynamicList[T]{}
	l.Push(elements...)
	return l
}

// DynamicList is a simple slice implementation of List[T].
// As a slice, it is dynamically-sized and not thread-safe.
type DynamicList[T any] []T

// Length returns how many elements are in the DynamicList.
func (l *DynamicList[T]) Length() int {
	return len(l.Elements())
}

// IsEmpty returns true if there are *no* Elements stored in the DynamicList.
func (l *DynamicList[T]) IsEmpty() bool {
	return l.Length() == 0
}

// IsNotEmpty returns true if there are Elements stored in the DynamicList.
func (l *DynamicList[T]) IsNotEmpty() bool {
	return !l.IsEmpty()
}

// At returns the pointer of the element at the given index from the DynamicList.
// If there is no element at the given index, nil will be returned.
func (l *DynamicList[T]) At(i int) (at *T) {
	if l.Length()-1 < i || i < 0 {
		return
	}
	at = &l.Elements()[i]
	return
}

// ElementAt returns the element at the given index from the DynamicList.
// If there is no element at the given index, panics.
func (l *DynamicList[T]) ElementAt(i int) T {
	return l.Elements()[i]
}

// Elements returns a built-in slice with all elements in the DynamicList.
func (l *DynamicList[T]) Elements() []T {
	return *l
}

// Push add the given elements in the DynamicList, and then returns itself.
func (l *DynamicList[T]) Push(elements ...T) List[T] {
	*l = append(l.Elements(), elements...)
	return l
}

// Clone returns an identical DynamicList from the original.
func (l *DynamicList[T]) Clone() List[T] {
	return NewDynamicList(l.Elements()...)
}

// FirstElement returns the first element in the DynamicList.
// If List is empty (see IsEmpty), panics
func (l *DynamicList[T]) FirstElement() T {
	return l.ElementAt(0)
}

// First returns the pointer of the first element in the DynamicList.
// If DynamicList is empty (see IsEmpty), nil will be returned.
func (l *DynamicList[T]) First() *T {
	return l.At(0)
}

// LastElement returns the last element in the List.
// If List is empty (see IsEmpty), panics.
func (l *DynamicList[T]) LastElement() T {
	return l.ElementAt(l.Length() - 1)
}

// Last returns the pointer of the last element in the List.
// If List is empty (see IsEmpty), nil will be returned.
func (l *DynamicList[T]) Last() *T {
	return l.At(l.Length() - 1)
}

// FirstIndexWhere returns the index of the first element which satisfies the predicate.
// If no element satisfies the predicate, -1 will be returned.
func (l *DynamicList[T]) FirstIndexWhere(handler Predicate[T]) int {
	for i, v := range l.Elements() {
		if handler(v) {
			return i
		}
	}
	return -1
}

// FirstWhere returns the pointer of the first element which satisfies the predicate.
// If no element satisfies the predicate, nil will be returned.
func (l *DynamicList[T]) FirstWhere(handler Predicate[T]) *T {
	for i, v := range l.Elements() {
		if handler(v) {
			return l.At(i)
		}
	}
	return nil
}

// FirstElementWhere returns the first element which satisfies the predicate.
// If no element satisfies the predicate, panics.
func (l *DynamicList[T]) FirstElementWhere(handler Predicate[T]) T {
	for _, v := range l.Elements() {
		if handler(v) {
			return v
		}
	}
	panic("no element satisfies the predicate")
}

// LastIndexWhere returns the index of the last element which satisfies the predicate.
// If no element satisfies the predicate, -1 will be returned.
func (l *DynamicList[T]) LastIndexWhere(handler Predicate[T]) int {
	idx := -1
	for i, v := range l.Elements() {
		if handler(v) {
			idx = i
		}
	}
	return idx
}

// LastWhere returns the pointer of the last element which satisfies the predicate.
// If no element satisfies the predicate, nil will be returned.
func (l *DynamicList[T]) LastWhere(handler Predicate[T]) *T {
	var e *T
	for i, v := range l.Elements() {
		if handler(v) {
			e = l.At(i)
		}
	}
	return e
}

// LastElementWhere returns the last element which satisfies the predicate.
// If no element satisfies the predicate, panics.
func (l *DynamicList[T]) LastElementWhere(handler Predicate[T]) T {
	var e T
	found := false
	for _, v := range l.Elements() {
		if handler(v) {
			e = v
			found = true
		}
	}
	if !found {
		panic("no element satisfies the predicate")
	}
	return e
}

// IndexWhere returns a List[int] for all element index which satisfies the predicate.
// If no element satisfies the predicate, an empty List will be returned.
func (l *DynamicList[T]) IndexWhere(handler Predicate[T]) List[int] {
	list := NewDynamicList[int]()
	for i, v := range l.Elements() {
		if handler(v) {
			list.Push(i)
		}
	}
	return list
}

// Where returns a List with all the elements which satisfies the predicate.
// If no element satisfies the predicate, an empty List will be returned.
func (l *DynamicList[T]) Where(handler Predicate[T]) List[T] {
	selected := NewDynamicList[T]()
	for _, v := range l.Elements() {
		if handler(v) {
			selected.Push(v)
		}
	}
	return selected
}

// Map iterates over the element of the List calling Mapper, and return a new List with the results.
func (l *DynamicList[T]) Map(handler Mapper[T]) List[any] {
	mapped := NewDynamicList[any]()
	for _, v := range l.Elements() {
		mapped.Push(handler(v))
	}
	return mapped
}

// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be the accumulator for the next.
// The final result will be returned.
func (l *DynamicList[T]) Reduce(reducer Reducer[T], accumulator any) any {
	for i, v := range l.Elements() {
		accumulator = reducer(accumulator, v, i)
	}
	return accumulator
}

// Every returns true if every element in the List satisfies the predicate.
func (l *DynamicList[T]) Every(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if !handler(v) {
			return false
		}
	}
	return true
}

// Some returns true if at least one element in the List satisfies the predicate.
func (l *DynamicList[T]) Some(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if handler(v) {
			return true
		}
	}
	return false
}

// None returns true no element in the List satisfy the predicate.
func (l *DynamicList[T]) None(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if handler(v) {
			return false
		}
	}
	return true
}

// Pop removes the last element from the List and returns itself.
func (l *DynamicList[T]) Pop() List[T] {
	*l = l.Elements()[0 : l.Length()-1]
	return l
}

// Shift removes the first element from the List and then returns itself.
func (l *DynamicList[T]) Shift() List[T] {
	*l = l.Elements()[1:l.Length()]
	return l
}

// Set sets the given element at the given index, and then returns itself.
func (l *DynamicList[T]) Set(index int, element T) List[T] {
	l.ElementAt(index)
	at := l.At(index)
	*at = element
	return l
}

// Interval returns a new List with all elements between the *from* and *to* indexes.
func (l *DynamicList[T]) Interval(from, to int) List[T] {
	return NewDynamicList[T](l.Elements()[from : to+1]...)
}

// String returns a string representation of the List.
func (l *DynamicList[T]) String() string {
	return fmt.Sprint(l.Elements())
}

// Sort receives a Sorter function to sort its elements, and returns itself after sorted.
func (l *DynamicList[T]) Sort(sorter Sorter[T]) List[T] {
	changed := true
	for changed {
		changed = l.sort(sorter)
	}
	return l
}

// Clear removes all elements from the List, making it empty, and then returns itself.
func (l *DynamicList[T]) Clear() List[T] {
	*l = []T{}
	return l
}

func (l *DynamicList[T]) sort(sorter Sorter[T]) bool {
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
