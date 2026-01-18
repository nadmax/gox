// Package set provides a generic Set data structure.
package set

// Set is a generic set implementation using a map.
type Set[T comparable] struct {
	items map[T]struct{}
}

// New creates and returns a new empty Set.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// Of creates a new Set containing the provided elements.
func Of[T comparable](elements ...T) *Set[T] {
	s := New[T]()
	for _, elem := range elements {
		s.Add(elem)
	}
	return s
}

// Add inserts an element into the set.
func (s *Set[T]) Add(element T) {
	s.items[element] = struct{}{}
}

// Remove deletes an element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.items, element)
}

// Contains checks if an element exists in the set.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.items[element]
	return exists
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.items)
}

// IsEmpty returns true if the set has no elements.
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// ToSlice returns a slice containing all elements in the set.
// The order of elements is not guaranteed.
func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.items))
	for elem := range s.items {
		result = append(result, elem)
	}
	return result
}

// Clone creates a shallow copy of the set.
func (s *Set[T]) Clone() *Set[T] {
	clone := New[T]()
	for elem := range s.items {
		clone.Add(elem)
	}
	return clone
}

// Union returns a new set containing all elements from both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := s.Clone()
	for elem := range other.items {
		result.Add(elem)
	}
	return result
}

// Intersection returns a new set containing only elements present in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	// Iterate over the smaller set for efficiency
	smaller, larger := s, other
	if other.Len() < s.Len() {
		smaller, larger = other, s
	}
	for elem := range smaller.items {
		if larger.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Difference returns a new set containing elements in s but not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	for elem := range s.items {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// SymmetricDifference returns a new set containing elements in either set but not in both.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	result := New[T]()
	for elem := range s.items {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	for elem := range other.items {
		if !s.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// IsSubset returns true if all elements of s are in other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}
	for elem := range s.items {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s contains all elements of other.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// Equal returns true if both sets contain exactly the same elements.
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for elem := range s.items {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// ForEach iterates over all elements in the set and applies the given function.
func (s *Set[T]) ForEach(fn func(T)) {
	for elem := range s.items {
		fn(elem)
	}
}

// Filter returns a new set containing only elements that satisfy the predicate.
func (s *Set[T]) Filter(predicate func(T) bool) *Set[T] {
	result := New[T]()
	for elem := range s.items {
		if predicate(elem) {
			result.Add(elem)
		}
	}
	return result
}
