package tsync

import (
	"iter"
	"sync"
)

// Map is a type-safe [sync.Map] that supports iterators.
type Map[K comparable, V any] struct{ m sync.Map }

// Clear deletes all the entries, resulting in an empty Map.
func (m *Map[K, V]) Clear() {
	m.m.Clear()
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// If there is no current value for key in the map, CompareAndDelete returns false.
func (m *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.m.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old.
func (m *Map[K, V]) CompareAndSwap(key K, old, nw V) (swapped bool) {
	return m.m.CompareAndSwap(key, old, nw)
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

// Load returns the value stored in the map for a key, or the zero value if no value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, false
	}
	return v.(V), true //nolint: errcheck,forcetypeassert // Known to be correct.
}

// Get is like Load but without the ok result.
func (m *Map[K, V]) Get(key K) V {
	v, _ := m.Load(key)
	return v
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, false
	}
	return v.(V), true //nolint: errcheck,forcetypeassert // Known to be correct.
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, ok := m.m.LoadOrStore(key, value)
	return v.(V), ok //nolint: errcheck,forcetypeassert // Known to be correct.
}

// All is like maps.All and uses (*sync.Map).Range under the hood:
//
// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's contents:
// no key will be visited more than once, but if the value for any key is stored
// or deleted concurrently (including by f), Range may reflect any mapping for
// that key from any point during the Range call. Range does not block other methods
// on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns false after a constant number of calls.
func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m.m.Range(func(key, value any) bool {
			return yield(key.(K), value.(V)) //nolint: errcheck,forcetypeassert // Known to be correct.
		})
	}
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, loaded := m.m.Swap(key, value)
	if !loaded {
		return previous, false
	}
	return v.(V), true //nolint: errcheck,forcetypeassert // Known to be correct.
}
