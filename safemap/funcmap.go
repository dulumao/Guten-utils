package safemap

import (
	"html/template"
	"Tada/src/components/util/conv"
)

func (self *SafeMap) ItemsToFuncMap() template.FuncMap {
	self.lock.RLock()
	defer self.lock.RUnlock()

	r := make(template.FuncMap)

	for k, v := range self.sm {
		r[conv.String(k)] = v
	}

	return r
}
