package set

// BytesIterationer 字节迭代器接口.
type BytesIterationer interface {
	// Next 是否有下一个数据.
	Next() bool
	// Key 获取数据.
	Key() []byte
	// Release 释放迭代.
	Release()
}
