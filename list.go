package iterable

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
	if l.Length()-1 <= i {
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
	return l.Elements()[0]
}

// First returns the pointer of the first element in the List
// if isEmpty, nil will be returned
func (l *List[T]) First() *T {
	return &l.Elements()[0]
}

// LastElement returns the last element in the List
// if isEmpty, panics
func (l *List[T]) LastElement() T {
	return l.Elements()[l.Length()-1]
}

// Last returns the pointer of the last element in the List
// if isEmpty, nil will be returned
func (l *List[T]) Last() *T {
	return &l.Elements()[l.Length()-1]
}

// FirstIndexWhere returns the index of the first element witch satisfies the predicate
// if no element satisfies the predicate, -1 will be returned
func (l *List[T]) FirstIndexWhere(handler PredicateHandler[T]) int {
	for i, v := range l.Elements() {
		if handler(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere returns the index of the last element witch satisfies the predicate
// if no element satisfies the predicate, -1 will be returned
func (l *List[T]) LastIndexWhere(handler PredicateHandler[T]) int {
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
func (l *List[T]) IndexWhere(handler PredicateHandler[T]) Iterable[int] {
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
func (l *List[T]) Where(handler PredicateHandler[T]) Iterable[T] {
	selected := NewList[T]()
	for _, v := range l.Elements() {
		if handler(v) {
			selected.Push(v)
		}
	}
	return selected
}
