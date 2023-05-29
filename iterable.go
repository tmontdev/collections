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
}
