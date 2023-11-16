package collections

// List is a collection interface to manage multiple values by int index keys (like built-in arrays and slices)
// It provides helper methods to easily handle data. The default implementation is *DynamicList
type List[T any] interface {
	// Length method returns how many element are in the List
	Length() int

	// IsEmpty returns true there are *no* Elements stored in the List
	IsEmpty() bool

	// IsNotEmpty returns true if there are Elements stored in the List
	IsNotEmpty() bool

	// At returns the pointer from the element in the given index from the List
	// If there is no element in the given index, nil will be returned
	At(int) *T

	// ElementAt returns the element in the given index from the List
	// If there is no element in the given index, panics
	ElementAt(int) T

	// Elements returns a built-in slice with all elements in the List
	Elements() []T

	// Push add the given elements in the List, and then returns itself
	Push(...T) List[T]

	// Clone returns an identical List from the original
	Clone() List[T]

	// FirstElement returns the first element in the List
	// if isEmpty, panics
	FirstElement() T

	// First returns the pointer of the first element in the List
	// if isEmpty, nil will be returned
	First() *T

	// LastElement returns the last element in the List
	// if isEmpty, panics
	LastElement() T

	// Last returns the pointer of the last element in the List
	// if isEmpty, nil will be returned
	Last() *T

	// FirstIndexWhere returns the index of the first element witch satisfies the predicate
	// if no element satisfies the predicate, -1 will be returned
	FirstIndexWhere(handler Predicate[T]) int

	// LastIndexWhere returns the index of the last element witch satisfies the predicate
	// if no element satisfies the predicate, -1 will be returned
	LastIndexWhere(handler Predicate[T]) int

	// IndexWhere returns a List[int] for all element index witch satisfies the predicate
	// if no element satisfies the predicate, an empty List will be returned
	IndexWhere(handler Predicate[T]) List[int]

	// Where returns a List with all the elements witch satisfies the predicate.
	// if no element satisfies the predicate, an empty List will be returned
	Where(handler Predicate[T]) List[T]

	// Map iterates over the element of the List calling Mapper, and return a new List with the results.
	Map(handler Mapper[T]) List[any]

	// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be accumulator for the next
	// The final result will be returned
	Reduce(reducer Reducer[T], accumulator any) any

	// Every returns true if every element in the List satisfy the predicate
	Every(handler Predicate[T]) bool

	// Some returns true if at least one element in the List satisfy the predicate
	Some(handler Predicate[T]) bool

	// None returns true no element in the List satisfy the predicate
	None(handler Predicate[T]) bool

	// Pop removes the last element from the List and returns itself
	Pop() List[T]

	// Shift removes the first element from the List and then returns itself
	Shift() List[T]

	// Set sets the given element in the given index, and then returns itself
	Set(index int, element T) List[T]

	// Interval returns a new List with all elements between the *from* and *to* indexes
	Interval(from, to int) List[T]

	// String returns a string representation of the List
	String() string

	// Sort receives a Sorter function to sort its elements, and returns itself after sorted
	Sort(sorter Sorter[T]) List[T]

	// Clear removes all elements from the List, making it empty, and then returns itself.
	Clear() List[T]
}
