package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDBSet(t *testing.T) {
	s := NewDBSet()
	defer s.Release()
	assert.NotEmpty(t, s.dir, "数据目录非空")

	test := []byte("test")
	str := []byte("str")
	err := []byte("err")
	err2 := []byte("err2")
	s.Add(test, str)

	assert.True(t, s.Has(test), "包含")
	assert.False(t, s.Has(err), "err不包含")
	assert.True(t, s.Has(str, test), "包含")
	assert.False(t, s.Has(test, err), "不包含")

	assert.True(t, s.HasAny(test, err), "包含任意一个")
	assert.False(t, s.HasAny(err2, err), "不包含任意一个")

	assert.False(t, s.Remove(test).Has(test), "删除后，不包含")

	assert.False(t, s.Empty(), "不为空")
	assert.Equal(t, s.Size(), 1, "数量")
	assert.True(t, s.Clear().Empty(), "为空")
}

func TestDBSet_Union(t *testing.T) {
	str1 := []byte("str1")
	str2 := []byte("str2")
	o := NewDBSet().Add(str2)
	defer o.Release()
	s := NewDBSet().Add(str1).Union(o)
	defer s.Release()

	assert.True(t, s.Has(str1, str2), "包含")
}

func TestDBSet_Intersect(t *testing.T) {
	str := []byte("str")
	str1 := []byte("str1")
	str2 := []byte("str2")
	o := NewDBSet().Add(str2, str)
	defer o.Release()

	s := NewDBSet().Add(str1, str).Intersect(o)
	defer s.Release()

	assert.True(t, s.Has(str), "包含")
	assert.Equal(t, s.Size(), 1, "1")
}

func TestDBSet_Complement(t *testing.T) {
	str := []byte("str")
	str1 := []byte("str1")
	str2 := []byte("str2")

	o := NewDBSet().Add(str2, str)
	defer o.Release()
	s := NewDBSet().Add(str1, str).Complement(o)
	defer s.Release()

	assert.True(t, s.Has(str2), "包含")
	assert.Equal(t, s.Size(), 1, "1")
}

func TestDBSet_Iteration(t *testing.T) {
	str1 := []byte("str1")
	str2 := []byte("str2")
	s := NewDBSet().Add(str1, str2)
	defer s.Release()

	i := s.Iteration()
	for i.Next() {
		assert.True(t, s.Has(i.Key()), "包含")
	}
}

func TestDBSet_Dir(t *testing.T) {
	s := NewDBSet().Dir("/tmp/a")
	defer s.Release()
	assert.Equal(t, s.dir, "/tmp/a")
}
