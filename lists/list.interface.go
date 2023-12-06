package lists

// IList is an interface which provides helper methods to easily handle arrays and slices.
// Each implementation may have specific behaviors. The default implementation is *List
type IList[T any] interface {
	// Length returns how many Elements are in the IList.
	Length() int

	// IsEmpty returns true if there are *no* Elements stored in the IList.
	IsEmpty() bool

	// IsNotEmpty returns true if there are Elements stored in the IList.
	IsNotEmpty() bool

	// At returns the pointer of the element at the given index from the IList.
	// If there is no element at the given index, nil will be returned.
	At(int) *T

	// ElementAt returns the element at the given index from the IList.
	// If there is no element at the given index, panics.
	ElementAt(int) T

	// Elements returns a built-in slice with all elements in the IList.
	Elements() []T

	// Push add the given elements in the IList, and then returns itself.
	Push(...T) IList[T]

	// Clone returns an identical IList from the original.
	Clone() IList[T]

	// FirstElement returns the first element in the IList.
	// If IList is empty (see IsEmpty), panics
	FirstElement() T

	// First returns the pointer of the first element in the IList.
	// If IList is empty (see IsEmpty), nil will be returned.
	First() *T

	// LastElement returns the last element in the IList.
	// If IList is empty (see IsEmpty), panics.
	LastElement() T

	// Last returns the pointer of the last element in the IList.
	// If IList is empty (see IsEmpty), nil will be returned.
	Last() *T

	// FirstIndexWhere returns the index of the first element which satisfies the predicate.
	// If no element satisfies the predicate, -1 will be returned.
	FirstIndexWhere(handler Predicate[T]) int

	// FirstWhere returns the pointer of the first element which satisfies the predicate.
	// If no element satisfies the predicate, nil will be returned.
	FirstWhere(handler Predicate[T]) *T

	// FirstElementWhere returns the first element which satisfies the predicate.
	// If no element satisfies the predicate, panics.
	FirstElementWhere(handler Predicate[T]) T

	// LastIndexWhere returns the index of the last element which satisfies the predicate.
	// If no element satisfies the predicate, -1 will be returned.
	LastIndexWhere(handler Predicate[T]) int

	// LastWhere returns the pointer of the last element which satisfies the predicate.
	// If no element satisfies the predicate, nil will be returned.
	LastWhere(handler Predicate[T]) *T

	// LastElementWhere returns the last element which satisfies the predicate.
	// If no element satisfies the predicate, panics.
	LastElementWhere(handler Predicate[T]) T

	// IndexWhere returns a IList[int] for all element index which satisfies the predicate.
	// If no element satisfies the predicate, an empty IList will be returned.
	IndexWhere(handler Predicate[T]) IList[int]

	// Where returns a IList with all the elements which satisfies the predicate.
	// If no element satisfies the predicate, an empty IList will be returned.
	Where(handler Predicate[T]) IList[T]

	// Map iterates over the elements of the IList calling Mapper, and return a new IList with the results.
	Map(handler Mapper[T]) IList[any]

	// Reduce executes the Reducer for each element from the list with the given accumulator, and each result will be the accumulator for the next.
	// The final result will be returned.
	Reduce(reducer Reducer[T], accumulator any) any

	// Every returns true if every element in the IList satisfies the predicate.
	Every(handler Predicate[T]) bool

	// Some returns true if at least one element in the IList satisfies the predicate.
	Some(handler Predicate[T]) bool

	// None returns true no element in the IList satisfy the predicate.
	None(handler Predicate[T]) bool

	// Pop removes the last element from the IList and returns itself.
	Pop() IList[T]

	// Shift removes the first element from the IList and then returns itself.
	Shift() IList[T]

	// Set sets the given element at the given index, and then returns itself.
	Set(index int, element T) IList[T]

	// Interval returns a new IList with all elements between the *from* and *to* indexes.
	Interval(from, to int) IList[T]

	// String returns a string representation of the IList.
	String() string

	// Join returns the string representation of each element in the IList, separated by the given separator
	Join(separator string) string

	// Sort receives a Sorter function to sort its elements, and returns itself after sorted.
	Sort(sorter Sorter[T]) IList[T]

	// Clear removes all elements from the IList, making it empty, and then returns itself.
	Clear() IList[T]

	//IsDynamicallySized returns true if the IList implementation is dynamically-sized
	IsDynamicallySized() bool

	//IsThreadSafe returns true if the IList implementation is dynamically-sized
	IsThreadSafe() bool
}
