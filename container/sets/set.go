package sets

// Set is a set of elements
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet returns a set of elements with assigned type
func NewSet[T comparable](es ...T) *Set[T] {
	s := &Set[T]{
		m: make(map[T]struct{}),
	}
	for _, e := range es {
		s.Add(e)
	}
	return s
}

// Add elements to set s, if element is already in s this has no effect
func (s *Set[T]) Add(es ...T) {
	for _, e := range es {
		s.m[e] = struct{}{}
	}
}

// Remove elements from set s, if element is not in s this has no effect
func (s *Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(s.m, e)
	}
}

// Has report whether v is in s
func (s *Set[T]) Has(es ...T) bool {
	for _, e := range es {
		if _, ok := s.m[e]; !ok {
			return false
		}
	}

	return true
}

// Len report the elements number of s
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Clone create a new set with the same elements as s
func (s *Set[T]) Clone() *Set[T] {
	ret := &Set[T]{
		m: make(map[T]struct{}),
	}
	for k := range s.m {
		ret.m[k] = struct{}{}
	}
	return ret
}

// ToSlice transform set to slice
func (s *Set[T]) ToSlice() []T {
	ret := make([]T, 0, len(s.m))

	for e := range s.m {
		ret = append(ret, e)
	}
	return ret
}
