package collections

// Map is a collections of key/value pairs, from which you retrieve a value using its associated key.
// It provides helper methods to easily handle data. The default implementation is HashMap
type Map[K comparable, V any] interface {
	// Length returns how many values are stored in the Map.
	Length() int

	// IsEmpty returns true if there are *no* value stored in the Map.
	IsEmpty() bool

	// IsNotEmpty returns true if there are values stored in the Map.
	IsNotEmpty() bool

	// Where returns a new Map containing only the key/value which satisfies de Predicate
	Where(predicate KeyValuePredicate[K, V]) Map[K, V]

	// RemoveWhere deletes all key/value which satisfies the Predicate, and then returns itself
	RemoveWhere(predicate KeyValuePredicate[K, V]) Map[K, V]

	// Some returns true if one or more key/value satisfies de Predicate
	Some(predicate KeyValuePredicate[K, V]) bool

	// None returns true if *no* key/value stored in the Map satisfies the predicate.
	None(predicate KeyValuePredicate[K, V]) bool

	// Every returns true if every key/value stored in the Map satisfies the predicate.
	Every(predicate KeyValuePredicate[K, V]) bool

	// Set sets the given value in the given key, and then returns itself
	Set(key K, value V) Map[K, V]

	// Get returns the value stored in the given key from the Map
	Get(key K) V

	// Access returns the value stored in the given key (if stored)
	Access(key K) (V, bool)

	// Clone returns a new Map with the same keys and values from the original
	Clone() Map[K, V]

	// Has returns true if the given key is filled.
	Has(key K) bool

	// Keys returns a SimpleList with all keys
	Keys() List[K]

	// Values returns a SimpleList with all values
	Values() List[V]

	// Merge sets all key/value pairs from the given map in itself.
	// If replace is false, ignore key conflicts
	Merge(source map[K]V, replace bool) Map[K, V]
}
