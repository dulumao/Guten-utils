package safemap

import (
	"Tada/src/components/util/conv"
)

func (self *SafeMap) GetString(k interface{}) string {
	self.lock.RLock()
	defer self.lock.RUnlock()

	if val, ok := self.sm[k]; ok {
		return conv.String(val)
	}

	return ""
}

func (m *SafeMap) ItemsToString(k interface{}) map[interface{}]string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	r := make(map[interface{}]string)

	for k, v := range m.sm {
		r[k] = conv.String(v)
	}

	return r
}
