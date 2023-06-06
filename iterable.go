package iterable

type Iterable[T any] interface {
	// Length method returns how many element are in the Iterable
	Length() int

	// IsEmpty returns true if Length is zero
	IsEmpty() bool

	// IsNotEmpty returns true if Length is not zero
	IsNotEmpty() bool

	// At returns the pointer of the element in the given index from the Iterable
	// If there is no element in the index, nil will be returned
	At(int) *T

	// ElementAt returns the element in the given index from the Iterable
	// If there is no element in the index, panics
	ElementAt(int) T

	// Elements returns a built-in slice with all elements in the Iterable
	Elements() []T

	// Push add the given elements in the Iterable, and returns itself
	Push(...T) Iterable[T]

	// Clone returns an identical Iterable from the original
	Clone() Iterable[T]

	// FirstElement returns the first element in the Iterable
	// if isEmpty, panics
	FirstElement() T

	// First returns the pointer of the first element in the Iterable
	// if isEmpty, nil will be returned
	First() *T

	// LastElement returns the last element in the Iterable
	// if isEmpty, panics
	LastElement() T

	// Last returns the pointer of the last element in the Iterable
	// if isEmpty, nil will be returned
	Last() *T

	// FirstIndexWhere returns the index of the first element witch satisfies the predicate
	// if no element satisfies the predicate, -1 will be returned
	FirstIndexWhere(handler PredicateHandler[T]) int

	// LastIndexWhere returns the index of the last element witch satisfies the predicate
	// if no element satisfies the predicate, -1 will be returned
	LastIndexWhere(handler PredicateHandler[T]) int

	// IndexWhere returns an Iterable[int] for all element index witch satisfies the predicate
	// if no element satisfies the predicate, an empty Iterable will be returned
	IndexWhere(handler PredicateHandler[T]) Iterable[int]

	// Where returns an Iterable with all the elements witch satisfies the predicate.
	// if no element satisfies the predicate, an empty Iterable will be returned
	Where(handler PredicateHandler[T]) Iterable[T]

	// Map iterates over the elements of the Iterable calling MapHandler, and return a new Iterable with the results.
	Map(handler MapHandler[T]) Iterable[any]

	// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be accumulator for the next
	// The final result will be returned
	Reduce(reducer Reducer[T], accumulator any) any

	// Every returns true if every element in the Iterable satisfy the predicate
	Every(handler PredicateHandler[T]) bool

	// Some returns true if at least one element in the Iterable satisfy the predicate
	Some(handler PredicateHandler[T]) bool

	// None returns true no element in the Iterable satisfy the predicate
	None(handler PredicateHandler[T]) bool

	// Pop removes the last element from the Iterable and returns itself
	Pop() Iterable[T]

	// Shift removes the first element from the Iterable and returns itself
	Shift() Iterable[T]
}
