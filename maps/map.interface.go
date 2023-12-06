package maps

import "github.com/tmontdev/collections/lists"

// IMap is a collections of key/value pairs, from which you retrieve a value using its associated key.
// It provides helper methods to easily handle data. The default implementation is Map
type IMap[K comparable, V any] interface {
	// Length returns how many values are stored in the IMap.
	Length() int

	// IsEmpty returns true if there are *no* value stored in the IMap.
	IsEmpty() bool

	// IsNotEmpty returns true if there are values stored in the IMap.
	IsNotEmpty() bool

	// Where returns a new IMap containing only the key/value which satisfies de Predicate
	Where(predicate Predicate[K, V]) IMap[K, V]

	// RemoveWhere deletes all key/value which satisfies the Predicate, and then returns itself
	RemoveWhere(predicate Predicate[K, V]) IMap[K, V]

	// Some returns true if one or more key/value satisfies de Predicate
	Some(predicate Predicate[K, V]) bool

	// None returns true if *no* key/value stored in the IMap satisfies the predicate.
	None(predicate Predicate[K, V]) bool

	// Every returns true if every key/value stored in the IMap satisfies the predicate.
	Every(predicate Predicate[K, V]) bool

	// Set sets the given value in the given key, and then returns itself
	Set(key K, value V) IMap[K, V]

	// Get returns the value stored in the given key from the IMap
	Get(key K) V

	// Access returns the value stored in the given key (if stored)
	Access(key K) (V, bool)

	// Clone returns a new IMap with the same keys and values from the original
	Clone() IMap[K, V]

	// Has returns true if the given key is filled.
	Has(key K) bool

	// Keys returns a List with all keys
	Keys() lists.IList[K]

	// Values returns a List with all values
	Values() lists.IList[V]

	// Complement sets missing key/value pairs from the given map in itself.
	Complement(source IMap[K, V]) IMap[K, V]

	// SetFrom sets all key/value pairs from the given map in itself.
	SetFrom(source IMap[K, V]) IMap[K, V]

	Builtin() map[K]V

	HashMap() Map[K, V]

	Struct(str any) error
}
