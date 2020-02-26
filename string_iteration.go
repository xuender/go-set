package set

// StringIterationer 字符串迭代器接口.
type StringIterationer interface {
	// Next 是否有下一个字符串.
	Next() bool
	// Key 获取字符串.
	Key() string
}

// StringIteration 字符串迭代器.
type StringIteration struct {
	data  []string
	index int
}

// Next 是否有下一个字符串.
func (i *StringIteration) Next() bool {
	i.index++
	return i.index < len(i.data)
}

// Key 字符串.
func (i *StringIteration) Key() string {
	v := i.data[i.index]
	return v
}
