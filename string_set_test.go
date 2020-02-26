package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStringSet(t *testing.T) {
	s := NewStringSet()

	s.Add("test", "str")

	assert.True(t, s.Has("test"), "包含")
	assert.False(t, s.Has("err"), "err不包含")
	assert.True(t, s.Has("str", "test"), "包含")
	assert.False(t, s.Has("test", "err"), "不包含")

	assert.True(t, s.HasAny("test", "err"), "包含任意一个")
	assert.False(t, s.HasAny("err2", "err"), "不包含任意一个")

	assert.False(t, s.Remove("test").Has("test"), "删除后，不包含")

	assert.False(t, s.Empty(), "不为空")
	assert.True(t, s.Clear().Empty(), "为空")
}

func TestStringSet_Union(t *testing.T) {
	s := NewStringSet().Add("str1").Union(NewStringSet().Add("str2"))
	assert.True(t, s.Has("str1", "str2"), "包含")
}

func TestStringSet_Intersect(t *testing.T) {
	s := NewStringSet().Add("str1", "str").Intersect(NewStringSet().Add("str2", "str"))
	assert.True(t, s.Has("str"), "包含")
	assert.Equal(t, s.Size(), 1, "1")
}

func TestStringSet_Complement(t *testing.T) {
	s := NewStringSet().Add("str1", "str").Complement(NewStringSet().Add("str2", "str"))
	assert.True(t, s.Has("str2"), "包含")
	assert.Equal(t, s.Size(), 1, "1")
}

func TestStringSet_Iteration(t *testing.T) {
	s := NewStringSet().Add("str1", "str2")
	i := s.Iteration()
	for i.Next() {
		assert.True(t, s.Has(i.Key()), "包含")
	}
}
