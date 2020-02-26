package set

import (
	"os"
	"path"

	"github.com/lithammer/shortuuid"
	"github.com/syndtr/goleveldb/leveldb"
)

// DBSet 基于数据库的集合.
type DBSet struct {
	db  *leveldb.DB // 数据库
	dir string      // 数据目录
	id  string      // 数据库标识
	err error
}

// Release 释放资源.
func (s *DBSet) Release() *DBSet {
	if s.db != nil {
		s.db.Close()
		os.RemoveAll(path.Join(s.dir, s.id))
		s.db = nil
	}
	return s
}

// Add adds some bytes and returns set.
func (s *DBSet) Add(bytes ...[]byte) *DBSet {
	if len(bytes) == 0 {
		return s
	}
	s.init()
	e := []byte{}
	for _, bs := range bytes {
		if err := s.db.Put(bs, e, nil); err != nil {
			s.err = err
		}
	}
	return s
}

// Has 全部包含.
func (s *DBSet) Has(bytes ...[]byte) bool {
	if len(bytes) == 0 {
		return true
	}
	s.init()
	for _, bs := range bytes {
		ok, err := s.db.Has(bs, nil)
		if err != nil {
			s.err = err
			return false
		}
		if !ok {
			return false
		}
	}
	return true
}

// HasAny has any one.
func (s *DBSet) HasAny(bytes ...[]byte) bool {
	if len(bytes) == 0 {
		return true
	}
	s.init()
	for _, bs := range bytes {
		ok, err := s.db.Has(bs, nil)
		if err != nil {
			s.err = err
			return false
		}
		if ok {
			return true
		}
	}
	return false
}

// Remove 删除.
func (s *DBSet) Remove(bytes ...[]byte) *DBSet {
	if len(bytes) == 0 {
		return s
	}
	s.init()
	for _, bs := range bytes {
		if err := s.db.Delete(bs, nil); err != nil {
			s.err = err
		}
	}
	return s
}

// Empty 是否为空.
func (s *DBSet) Empty() bool {
	s.init()
	i := s.db.NewIterator(nil, nil)
	defer i.Release()
	return !i.Next()
}

// Size 尺寸.
func (s *DBSet) Size() int {
	s.init()
	i := s.db.NewIterator(nil, nil)
	defer i.Release()
	ret := 0
	for i.Next() {
		ret++
	}
	return ret
}

// Clear 清空集合.
func (s *DBSet) Clear() *DBSet {
	return s.Release()
}

// Dir 设置目录.
func (s *DBSet) Dir(dir string) *DBSet {
	if s.dir != dir {
		s.Release()
		s.dir = dir
	}
	return s
}

func (s *DBSet) init() {
	if s.db == nil {
		id := shortuuid.New()
		db, err := leveldb.OpenFile(path.Join(s.dir, id), nil)
		if err != nil {
			s.err = err
		}
		s.id = id
		s.db = db
	}
}

// Iteration 迭代.
func (s *DBSet) Iteration() BytesIterationer {
	return s.db.NewIterator(nil, nil)
}

// Union 并集.
func (s *DBSet) Union(other *DBSet) *DBSet {
	i := other.Iteration()
	defer i.Release()

	e := []byte{}
	for i.Next() {
		if err := s.db.Put(i.Key(), e, nil); err != nil {
			s.err = err
		}
	}
	return s
}

// Intersect 交集.
func (s *DBSet) Intersect(other *DBSet) *DBSet {
	i := s.Iteration()
	defer i.Release()

	for i.Next() {
		if !other.Has(i.Key()) {
			s.Remove(i.Key())
		}
	}
	return s
}

// Complement 补集.
func (s *DBSet) Complement(full *DBSet) *DBSet {
	s.Intersect(full)
	i := full.Iteration()
	defer i.Release()

	for i.Next() {
		if s.Has(i.Key()) {
			s.Remove(i.Key())
		} else {
			s.Add(i.Key())
		}
	}
	return s
}

// NewDBSet new databse set.
func NewDBSet() *DBSet {
	return &DBSet{dir: "/tmp"}
}
