package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	a := NewSet[string]("aaa", "bbb", "ccc")
	b := NewSet[string]("aaa", "ccc")
	c := NewSet[string]("aaa", "bbb")

	d := Intersect(a, b, c)
	assert.ElementsMatch(t, []string{"aaa"}, d.ToSlice())
}

func TestUnion(t *testing.T) {
	a := NewSet[string]("aaa", "bbb")
	b := NewSet[string]("aaa", "ccc")
	c := NewSet[string]("ccc", "bbb")

	d := Union(a, b, c)
	assert.ElementsMatch(t, []string{"aaa", "bbb", "ccc"}, d.ToSlice())
}

func TestDifference(t *testing.T) {
	a := NewSet[string]("aaa", "bbb", "ccc")
	b := NewSet[string]("aaa", "ccc")

	d := Difference(a, b)
	assert.ElementsMatch(t, []string{"bbb"}, d.ToSlice())
}

func TestSymmetricDifference(t *testing.T) {
	a := NewSet[string]("aaa", "bbb")
	b := NewSet[string]("bbb", "ccc")

	d := SymmetricDifference(a, b)
	assert.ElementsMatch(t, []string{"aaa", "ccc"}, d.ToSlice())
}
