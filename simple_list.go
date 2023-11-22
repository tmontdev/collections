package collections

import (
	"fmt"
)

// NewList returns a new SimpleList with the given elements
func NewList[T any](elements ...T) *SimpleList[T] {
	l := &SimpleList[T]{}
	l.Push(elements...)
	return l
}

// NewListFrom returns a new SimpleList with the given slice
func NewListFrom[T any](elements []T) *SimpleList[T] {
	l := &SimpleList[T]{}
	l.Push(elements...)
	return l
}

// SimpleList is a dynamically-sized and thread-unsafe implementation of List.
type SimpleList[T any] []T

// Length returns how many elements are in the SimpleList.
func (l *SimpleList[T]) Length() int {
	return len(l.Elements())
}

// IsEmpty returns true if there are *no* Elements stored in the SimpleList.
func (l *SimpleList[T]) IsEmpty() bool {
	return l.Length() == 0
}

// IsNotEmpty returns true if there are Elements stored in the SimpleList.
func (l *SimpleList[T]) IsNotEmpty() bool {
	return !l.IsEmpty()
}

// At returns the pointer of the element at the given index from the SimpleList.
// If there is no element at the given index, nil will be returned.
func (l *SimpleList[T]) At(i int) (at *T) {
	if l.Length()-1 < i || i < 0 {
		return
	}
	at = &l.Elements()[i]
	return
}

// ElementAt returns the element at the given index from the SimpleList.
// If there is no element at the given index, panics.
func (l *SimpleList[T]) ElementAt(i int) T {
	return l.Elements()[i]
}

// Elements returns a built-in slice with all elements in the SimpleList.
func (l *SimpleList[T]) Elements() []T {
	return *l
}

// Push add the given elements in the SimpleList, and then returns itself.
func (l *SimpleList[T]) Push(elements ...T) List[T] {
	*l = append(l.Elements(), elements...)
	return l
}

// Clone returns an identical SimpleList from the original.
func (l *SimpleList[T]) Clone() List[T] {
	return NewList(l.Elements()...)
}

// FirstElement returns the first element in the SimpleList.
// If SimpleList is empty (see IsEmpty), panics
func (l *SimpleList[T]) FirstElement() T {
	return l.ElementAt(0)
}

// First returns the pointer of the first element in the SimpleList.
// If SimpleList is empty (see IsEmpty), nil will be returned.
func (l *SimpleList[T]) First() *T {
	return l.At(0)
}

// LastElement returns the last element in the SimpleList.
// If SimpleList is empty (see IsEmpty), panics.
func (l *SimpleList[T]) LastElement() T {
	return l.ElementAt(l.Length() - 1)
}

// Last returns the pointer of the last element in the List.
// If List is empty (see IsEmpty), nil will be returned.
func (l *SimpleList[T]) Last() *T {
	return l.At(l.Length() - 1)
}

// FirstIndexWhere returns the index of the first element which satisfies the predicate.
// If no element satisfies the predicate, -1 will be returned.
func (l *SimpleList[T]) FirstIndexWhere(handler Predicate[T]) int {
	for i, v := range l.Elements() {
		if handler(v) {
			return i
		}
	}
	return -1
}

// FirstWhere returns the pointer of the first element which satisfies the predicate.
// If no element satisfies the predicate, nil will be returned.
func (l *SimpleList[T]) FirstWhere(handler Predicate[T]) *T {
	for i, v := range l.Elements() {
		if handler(v) {
			return l.At(i)
		}
	}
	return nil
}

// FirstElementWhere returns the first element which satisfies the predicate.
// If no element satisfies the predicate, panics.
func (l *SimpleList[T]) FirstElementWhere(handler Predicate[T]) T {
	for _, v := range l.Elements() {
		if handler(v) {
			return v
		}
	}
	panic("no element satisfies the predicate")
}

// LastIndexWhere returns the index of the last element which satisfies the predicate.
// If no element satisfies the predicate, -1 will be returned.
func (l *SimpleList[T]) LastIndexWhere(handler Predicate[T]) int {
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
func (l *SimpleList[T]) LastWhere(handler Predicate[T]) *T {
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
func (l *SimpleList[T]) LastElementWhere(handler Predicate[T]) T {
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

// IndexWhere returns a SimpleList[int] for all element index which satisfies the predicate.
// If no element satisfies the predicate, an empty SimpleList will be returned.
func (l *SimpleList[T]) IndexWhere(handler Predicate[T]) List[int] {
	list := NewList[int]()
	for i, v := range l.Elements() {
		if handler(v) {
			list.Push(i)
		}
	}
	return list
}

// Where returns a SimpleList with all the elements which satisfies the predicate.
// If no element satisfies the predicate, an empty SimpleList will be returned.
func (l *SimpleList[T]) Where(handler Predicate[T]) List[T] {
	selected := NewList[T]()
	for _, v := range l.Elements() {
		if handler(v) {
			selected.Push(v)
		}
	}
	return selected
}

// Map iterates over the element of the SimpleList calling Mapper, and return a new SimpleList with the results.
func (l *SimpleList[T]) Map(handler Mapper[T]) List[any] {
	mapped := NewList[any]()
	for _, v := range l.Elements() {
		mapped.Push(handler(v))
	}
	return mapped
}

// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be the accumulator for the next.
// The final result will be returned.
func (l *SimpleList[T]) Reduce(reducer Reducer[T], accumulator any) any {
	for i, v := range l.Elements() {
		accumulator = reducer(accumulator, v, i)
	}
	return accumulator
}

// Every returns true if every element in the List satisfies the predicate.
func (l *SimpleList[T]) Every(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if !handler(v) {
			return false
		}
	}
	return true
}

// Some returns true if at least one element in the List satisfies the predicate.
func (l *SimpleList[T]) Some(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if handler(v) {
			return true
		}
	}
	return false
}

// None returns true no element in the List satisfy the predicate.
func (l *SimpleList[T]) None(handler Predicate[T]) bool {
	for _, v := range l.Elements() {
		if handler(v) {
			return false
		}
	}
	return true
}

// Pop removes the last element from the List and returns itself.
func (l *SimpleList[T]) Pop() List[T] {
	*l = l.Elements()[0 : l.Length()-1]
	return l
}

// Shift removes the first element from the List and then returns itself.
func (l *SimpleList[T]) Shift() List[T] {
	*l = l.Elements()[1:l.Length()]
	return l
}

// Set sets the given element at the given index, and then returns itself.
func (l *SimpleList[T]) Set(index int, element T) List[T] {
	l.ElementAt(index)
	at := l.At(index)
	*at = element
	return l
}

// Interval returns a new SimpleList with all elements between the *from* and *to* indexes.
func (l *SimpleList[T]) Interval(from, to int) List[T] {
	return NewList[T](l.Elements()[from : to+1]...)
}

// String returns a string representation of the SimpleList.
func (l *SimpleList[T]) String() string {
	return fmt.Sprint(l.Elements())
}

// Join returns the string representation of each element in the List, separated by the given separator
func (l *SimpleList[T]) Join(separator string) (joined string) {
	if l.IsEmpty() {
		return
	}
	for i, e := range l.Elements() {
		if i == 0 {
			joined += fmt.Sprintf("%v", e)
			continue
		}
		joined += fmt.Sprintf("%s%v", separator, e)
	}
	return
}

// Sort receives a Sorter function to sort its elements, and returns itself after sorted.
func (l *SimpleList[T]) Sort(sorter Sorter[T]) List[T] {
	changed := true
	for changed {
		changed = l.sort(sorter)
	}
	return l
}

// Clear removes all elements from the SimpleList, making it empty, and then returns itself.
func (l *SimpleList[T]) Clear() List[T] {
	*l = []T{}
	return l
}

// IsDynamicallySized returns true, as SafeListDynamicList is a dynamically-sized implementation of List
func (l *SimpleList[T]) IsDynamicallySized() bool {
	return true
}

// IsThreadSafe returns false, as SimpleList is not a thread-safe implementation of List
func (l *SimpleList[T]) IsThreadSafe() bool {
	return false
}

func (l *SimpleList[T]) sort(sorter Sorter[T]) bool {
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
