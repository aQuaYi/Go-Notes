# sync.Map 是线程安全的map

sync.Map是Golang 1.9 的新特性。具有以下特点
1. 并发安全
1. 增删查改的时间复杂度是 amortized-constant-time 。

初始化
```go
var m sync.Map
```

暴露出来的方法有
```go
// Store 往m中存储键值对
func (m *Map) Store(key, value interface{})

// Load 查找 key 对应的值，当key不存在时，interface{}为nil， bool为false
func (m *Map) Load(key interface{}) (interface{},bool)

// LoadOrStore returns the existing value for the key if present.
// key存在，就返回对应的value 和 true
// Otherwise, it stores and returns the given value.
// 否则的话，就储存参数中的键值对，并返回参数中的value 和 false
// The loaded result is true if the value was loaded, false if stored.
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)

// Delete 删除key及其对应的value
func (m *Map) Delete(key interface{}) 

// Range calls f sequentially for each key and value present in the map.
// Range 会连续地对key，value 调用 f 函数。
// If f returns false, range stops the iteration.
// 当 f 返回 false 时， Range 会停止迭代。
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
// Range 迭代的内容，并不对应于map的某个快照的固定内容。不会有key会不止一次被访问，但是，如果
// key 对应的值在迭代期间被存储或者删除， Range 会返回反射出，在 Range 期间任意一个时间点的
// 映射值。
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
// Range 的时间复杂度O(N)，N为map的元素个数，即使 f 在固定次数后，返回false
func (m *Map) Range(f func(key, value interface{}) bool)
```