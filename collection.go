package collection

type Collection[T any] interface {
	// Length method returns how many element are in the Collection
	Length() int

	// IsEmpty returns true if Length is zero
	IsEmpty() bool

	// IsNotEmpty returns true if Length is not zero
	IsNotEmpty() bool

	// At returns the pointer of the element in the given index from the Collection
	// If there is no element in the index, nil will be returned
	At(int) *T

	// ElementAt returns the element in the given index from the Collection
	// If there is no element in the index, panics
	ElementAt(int) T

	// Elements returns a built-in slice with all elements in the Collection
	Elements() []T

	// Push add the given elements in the Collection, and returns itself
	Push(...T) Collection[T]

	// Clone returns an identical Collection from the original
	Clone() Collection[T]

	// FirstElement returns the first element in the Collection
	// if isEmpty, panics
	FirstElement() T

	// First returns the pointer of the first element in the Collection
	// if isEmpty, nil will be returned
	First() *T

	// LastElement returns the last element in the Collection
	// if isEmpty, panics
	LastElement() T

	// Last returns the pointer of the last element in the Collection
	// if isEmpty, nil will be returned
	Last() *T

	// FirstIndexWhere returns the index of the first element witch satisfies the predicate
	// if no element satisfies the predicate, -1 will be returned
	FirstIndexWhere(handler Predicate[T]) int

	// LastIndexWhere returns the index of the last element witch satisfies the predicate
	// if no element satisfies the predicate, -1 will be returned
	LastIndexWhere(handler Predicate[T]) int

	// IndexWhere returns a Collection[int] for all element index witch satisfies the predicate
	// if no element satisfies the predicate, an empty Collection will be returned
	IndexWhere(handler Predicate[T]) Collection[int]

	// Where returns a Collection with all the elements witch satisfies the predicate.
	// if no element satisfies the predicate, an empty Collection will be returned
	Where(handler Predicate[T]) Collection[T]

	// Map iterates over the elements of the Collection calling Mapper, and return a new Collection with the results.
	Map(handler Mapper[T]) Collection[any]

	// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be accumulator for the next
	// The final result will be returned
	Reduce(reducer Reducer[T], accumulator any) any

	// Every returns true if every element in the Collection satisfy the predicate
	Every(handler Predicate[T]) bool

	// Some returns true if at least one element in the Collection satisfy the predicate
	Some(handler Predicate[T]) bool

	// None returns true no element in the Collection satisfy the predicate
	None(handler Predicate[T]) bool

	// Pop removes the last element from the Collection and returns itself
	Pop() Collection[T]

	// Shift removes the first element from the Collection and returns itself
	Shift() Collection[T]

	// Set sets the given element in the given index, and returns itself
	Set(index int, element T) Collection[T]

	// Interval returns a new Collection with all elements between from and to given indexes
	Interval(from, to int) Collection[T]

	// String returns a string representation of the Collection
	String() string

	// Sort receive a Sorter function to sort its elements, and returns itself after sorted
	Sort(sorter Sorter[T]) Collection[T]
}
