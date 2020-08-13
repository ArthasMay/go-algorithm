package simplelru

type LRUCache interface {
	// Add: 添加缓存
	Add(key, value interface{}) bool
	Get(key interface{}) (value interface{}, ok bool)
	Contains(key interface{}) (ok bool)
	Peek(key interface{}) (value interface{}, ok bool)
	Remove(key interface{}) bool
	RemoveOldest() (interface{}, interface{}, bool)
	GetOldest() (interface{}, interface{}, bool)
	Keys() []interface{}
	// Returns the number of items in the cache.
	Len() int

	// Clears all cache entries.
	Purge()

	// Resizes cache, returning number evicted
	Resize(int) int
}
