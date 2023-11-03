package collection

type Dictionary[K comparable, V any] interface {
	// Length method return how many elements are stored in Dictionary
	Length() int

	// IsEmpty method return true if there are no values stored in the Dictionary.
	IsEmpty() bool

	// IsNotEmpty method return true if there are one or more values stored in the Dictionary
	IsNotEmpty() bool

	// Where returns a new Dictionary containing only the elements witch satisfies de Predicate
	Where(predicate KeyPredicate[K, V]) Dictionary[K, V]

	// RemoveWhere deletes all elements witch satisfies the Predicate, and returns itself
	RemoveWhere(predicate KeyPredicate[K, V]) Dictionary[K, V]

	// Some returns true if one or more elements satisfies de Predicate
	Some(predicate KeyPredicate[K, V]) bool

	// None returns true when no element satisfies the Predicate
	None(predicate KeyPredicate[K, V]) bool

	// Every returns true when all the elements satisfies the Predicate
	Every(predicate KeyPredicate[K, V]) bool

	// Set sets the given value in the given key, and returns itself
	Set(key K, value V) Dictionary[K, V]

	// Get returns the value stored in the given key from the Dictionary
	Get(key K) V

	// Access returns the value stored in the given key (if stored)
	Access(key K) (V, bool)

	// Clone returns a new Dictionary with the same keys and values from the original
	Clone() Dictionary[K, V]

	// Has returns true if the given key is filled.
	Has(key K) bool

	// Keys returns a List with all
}
