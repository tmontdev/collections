package iterable

type Iterable[T any] interface {
	// Length method returns how many element are in the Iterable
	Length() int

	// IsEmpty returns true if Length is zero
	IsEmpty() bool

	// IsNotEmpty returns true if Length is not zero
	IsNotEmpty() bool

	// At returns the pointer of the element in the given index from the Iterable
	At(int) *T

	ElementAt(int) T

	Elements() []T
}
