package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	val, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(val)
		return val.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	v, ok := c.items[key]
	if ok {
		v.Value = value
		c.queue.MoveToFront(v)
	} else {
		item := c.queue.PushFront(value)
		item.Key = key
		c.items[key] = item
	}
	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		delete(c.items, c.queue.Back().Key)
	}
	return ok
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
