package set

// StringSet is string set.
type StringSet struct {
	data map[string]bool
}

// Add adds some string and returns set.
func (s *StringSet) Add(strings ...string) *StringSet {
	for _, str := range strings {
		s.data[str] = true
	}
	return s
}

// Union 并集.
func (s *StringSet) Union(other *StringSet) *StringSet {
	for str := range other.data {
		s.data[str] = true
	}
	return s
}

// Intersect 交集.
func (s *StringSet) Intersect(other *StringSet) *StringSet {
	for str := range s.data {
		if _, ok := other.data[str]; !ok {
			delete(s.data, str)
		}
	}
	return s
}

// Complement 补集.
func (s *StringSet) Complement(full *StringSet) *StringSet {
	for str := range s.data {
		if _, ok := full.data[str]; !ok {
			delete(s.data, str)
		}
	}
	for str := range full.data {
		if _, ok := s.data[str]; ok {
			delete(s.data, str)
		} else {
			s.data[str] = true
		}
	}
	return s
}

// Empty 是否为空.
func (s *StringSet) Empty() bool {
	return len(s.data) == 0
}

// Size 尺寸.
func (s *StringSet) Size() int {
	return len(s.data)
}

// Clear 清空集合.
func (s *StringSet) Clear() *StringSet {
	s.data = map[string]bool{}
	return s
}

// Has 全包含.
func (s *StringSet) Has(strings ...string) bool {
	for _, str := range strings {
		if _, ok := s.data[str]; !ok {
			return false
		}
	}
	return true
}

// HasAny has any one.
func (s *StringSet) HasAny(strings ...string) bool {
	for _, str := range strings {
		if _, ok := s.data[str]; ok {
			return true
		}
	}
	return false
}

// Remove 删除.
func (s *StringSet) Remove(strings ...string) *StringSet {
	for _, str := range strings {
		delete(s.data, str)
	}
	return s
}

// Iteration 迭代.
func (s *StringSet) Iteration() StringIterationer {
	j := 0
	data := make([]string, len(s.data))
	for k := range s.data {
		data[j] = k
		j++
	}
	return &StringIteration{data: data, index: -1}
}

// NewStringSet new string set.
func NewStringSet() *StringSet {
	return &StringSet{data: map[string]bool{}}
}
