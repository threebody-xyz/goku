package sets

// Intersect returns the intersection of some sets(eg. A, B), each element is in A and B
func Intersect[T comparable](sets ...*Set[T]) *Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	ret := sets[0].Clone()
	if len(sets) == 1 {
		return ret
	}

	for v := range ret.m {
		for i := 1; i < len(sets); i++ {
			if !sets[i].Has(v) {
				ret.Remove(v)
				break
			}
		}
	}
	return ret
}

// Union returns the union of some sets(eg. A, B), each element is in A or B
func Union[T comparable](sets ...*Set[T]) *Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}

	ret := NewSet[T]()
	for _, s := range sets {
		for v := range s.m {
			ret.Add(v)
		}
	}
	return ret
}

// Difference returns the set that is in A but not in B
func Difference[T comparable](a, b *Set[T]) *Set[T] {
	ret := NewSet[T]()
	for v := range a.m {
		if !b.Has(v) {
			ret.Add(v)
		}
	}
	return ret
}

// SymmetricDifference returns the set that belong to one set(A or B) but not to the other
// SymmetricDifference(a, b) = SymmetricDifference(b, a) = Union(Difference(a, b), Difference(b, a))
func SymmetricDifference[T comparable](a, b *Set[T]) *Set[T] {
	ret := NewSet[T]()
	for v := range a.m {
		if !b.Has(v) {
			ret.Add(v)
		}
	}
	for v := range b.m {
		if !a.Has(v) {
			ret.Add(v)
		}
	}
	return ret
}
