package safemap

import (
	"html/template"
)

func (self *SafeMap) GetTemplate(k interface{}) *template.Template {
	self.lock.RLock()
	defer self.lock.RUnlock()

	if val, ok := self.sm[k]; ok {
		if template, ok := val.(*template.Template); ok {
			return template
		}
	}

	return nil
}

func (self *SafeMap) ItemsToTemplate(k interface{}) map[interface{}]*template.Template {
	self.lock.RLock()
	defer self.lock.RUnlock()

	r := make(map[interface{}]*template.Template)

	for k, v := range self.sm {
		if template, ok := v.(*template.Template); ok {
			r[k] = template
		}
	}

	return r
}
