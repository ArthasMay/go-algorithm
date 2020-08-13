package simplelru

import (
	"container/list"
	"errors"
)

type EvictCallBack func(key interface{}, value interface{})

type LRU struct {
	size       int
	envictList *list.List
	items      map[interface{}]*list.Element
	onEvict    EvictCallBack
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewLRU(size int, onEvict EvictCallBack) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("Must provide a positive size")
	}

	c := &LRU{
		size:       size,
		envictList: list.New(),
		items:      make(map[interface{}]*list.Element),
		onEvict:    onEvict,
	}
	return c, nil
}

func (c *LRU) Purge() {

}

func (c *LRU) Add(key, value interface{}) (evicted bool) {
	// 检查LRU Cache中是否存key对应的缓存对象
	if ent, ok := c.items[key]; ok {
		c.envictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return false
	}

	// 添加新的item
	ent := &entry{key, value}
	entry := c.envictList.PushFront(ent)
	c.items[key] = entry

	evict := c.envictList.Len() > c.size
	if evict {

	}
	return evict
}

func (c *LRU) Get(key interface{}) (value interface{}, ok bool) {
	if ent, ok := c.items[key]; ok {
		c.envictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil {
			return nil, false
		}
		return ent.Value.(*entry).value, true
	}
	return
}

func (c *LRU) Contains(key interface{}) (ok bool) {
	_, ok = c.items[key]
	return ok
}

func (c *LRU) Peek(key interface{}) (value interface{}, ok bool) {
	var ent *list.Element
	if ent, ok = c.items[key]; ok {
		return ent.Value.(*entry).value, true
	}
	return nil, ok
}

func (c *LRU) Remove(key interface{}) (present bool) {
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

func (c *LRU) RemoveOldest() (key, value interface{}, ok bool) {
	ent := c.envictList.Back()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

func (c *LRU) removeOldest() {
	ent := c.envictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

func (c *LRU) GetOldest() (key, value interface{}, ok bool) {
	ent := c.envictList.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

func (c *LRU) Keys() []interface{} {
	keys := make([]interface{}, len(c.items))
	i := 0
	for ent := c.envictList.Back(); ent != nil; ent = ent.Prev() {
		keys[i] = ent.Value.(*entry).key
	}
	return keys
}

func (c *LRU) Len() int {
	return c.envictList.Len()
}

func (c *LRU) Resize(size int) (evicted int) {
	diff := c.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		c.removeOldest()
	}
	c.size = size
	return diff
}

func (c *LRU) removeElement(e *list.Element) {
	c.envictList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
	if c.onEvict != nil {
		c.onEvict(kv.key, kv.value)
	}
}
