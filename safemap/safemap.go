package safemap

import (
	"sync"
)

// SafeMap is a map with lock
type SafeMap struct {
	lock *sync.RWMutex
	sm   map[interface{}]interface{}
}

// NewSafeMap return new SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		sm:   make(map[interface{}]interface{}),
	}
}

// Get from maps return the k's value
func (self *SafeMap) Get(k interface{}) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()

	if val, ok := self.sm[k]; ok {
		return val
	}

	return nil
}

// Set Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (self *SafeMap) Set(k interface{}, v interface{}) bool {
	self.lock.Lock()
	defer self.lock.Unlock()

	if val, ok := self.sm[k]; !ok {
		self.sm[k] = v
	} else if val != v {
		self.sm[k] = v
	} else {
		return false
	}

	return true
}

// Check Returns true if k is exist in the map.
func (self *SafeMap) Check(k interface{}) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()

	_, ok := self.sm[k]

	return ok
}

// Delete the given key and value.
func (self *SafeMap) Delete(k interface{}) {
	self.lock.Lock()
	defer self.lock.Unlock()

	delete(self.sm, k)
}

// Items returns all items in safemap.
func (self *SafeMap) Items() map[interface{}]interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()

	r := make(map[interface{}]interface{})

	for k, v := range self.sm {
		r[k] = v
	}

	return r
}

// Count returns the number of items within the map.
func (self *SafeMap) Count() int {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return len(self.sm)
}

func (self *SafeMap) BatchSet(values map[string]string) {
	self.lock.Lock()
	defer self.lock.Unlock()

	for k, v := range values {
		self.sm[k] = v
		// if val, ok := self.sm[k]; !ok {
		// 	self.sm[k] = v
		// } else if val != v {
		// 	self.sm[k] = v
		// } else {
		// 	// 不处理
		// }
	}
}

// 批量删除键值对
func (self *SafeMap) BatchRemove(keys []string) {
	self.lock.Lock()
	defer self.lock.Unlock()

	for _, k := range keys {
		delete(self.sm, k)
	}
}

func (self *SafeMap) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.sm = make(map[interface{}]interface{})
}

// 哈希表是否为空
func (self *SafeMap) IsEmpty() bool {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return len(self.sm) == 0
}

// 是否存在某个键
func (self *SafeMap) Has(key string) bool {
	self.lock.RLock()
	defer self.lock.RLocker()

	_, exists := self.sm[key]

	return exists
}

func (self *SafeMap) Keys() []interface{} {
	self.lock.RLock()
	defer self.lock.RLocker()

	keys := make([]interface{}, 0)

	for key, _ := range self.sm {
		keys = append(keys, key)
	}

	return keys
}

// 返回值列表(注意是随机排序)
func (self *SafeMap) Values() []interface{} {
	self.lock.RLock()
	defer self.lock.RLocker()

	vals := make([]interface{}, 0)

	for _, val := range self.sm {
		vals = append(vals, val)
	}

	return vals
}

// 当键名存在时返回其键值，否则写入指定的键值
func (self *SafeMap) GetOrSet(key string, value string) (interface{}, bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	v, ok := self.sm[key]

	return v, ok
}

// 哈希表克隆
func (self *SafeMap) Clone() map[interface{}]interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()

	c := make(map[interface{}]interface{})

	for k, v := range self.sm {
		c[k] = v
	}

	return c
}
