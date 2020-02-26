package set

// Set 接口.
type Set interface {
	// Union 合并.
	Union(other Set) Set
	// Intersect 交集.
	Intersect(other Set) Set
	// Empty 集合是否为空.
	Empty() bool
	// Clear 清空.
	Clear() Set
}
