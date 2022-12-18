package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetString(t *testing.T) {
	s := NewSet[string]()
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, []string{}, s.ToSlice())

	s.Add("aaa", "bbb", "ccc")
	assert.ElementsMatch(t, []string{"aaa", "bbb", "ccc"}, s.ToSlice())
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, true, s.Has("aaa", "ccc"))

	s.Remove("bbb", "ddd")
	assert.ElementsMatch(t, []string{"aaa", "ccc"}, s.ToSlice())
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, true, s.Has("aaa"))
	assert.Equal(t, false, s.Has("aaa", "ddd"))

	c := s.Clone()
	c.Add("ddd")
	assert.ElementsMatch(t, []string{"aaa", "ccc"}, s.ToSlice())
	assert.ElementsMatch(t, []string{"aaa", "ccc", "ddd"}, c.ToSlice())
}

func TestSetCustomType(t *testing.T) {
	type User struct {
		name string
	}
	s := NewSet[User]()
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, []User{}, s.ToSlice())

	s.Add(User{"aaa"}, User{"bbb"}, User{"ccc"})
	assert.ElementsMatch(t, []User{{"aaa"}, {"bbb"}, {"ccc"}}, s.ToSlice())
	assert.Equal(t, 3, s.Len())
}
