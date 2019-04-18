package httplib

import (
	"bytes"
	"net/url"
)

type IParams interface {
	Set(interface{}, interface{})
	Size() int
	Keys() []interface{}
	Encode() string
	Get(interface{}) (interface{}, bool)
}
type Params struct {
	data map[string][]string
}

func NewParams() *Params {
	return &Params{
		data: map[string][]string{},
	}
}

func (self *Params) Set(key interface{}, value interface{}) {
	ks, _ := key.(string)
	vs, _ := value.(string)

	if params, ok := self.data[ks]; ok {
		self.data[ks] = append(params, vs)
	} else {
		self.data[ks] = []string{vs}
	}
}

func (self *Params) Size() int {
	return len(self.data)
}

func (self *Params) Keys() []interface{} {
	keys := make([]interface{}, 0)

	for k := range self.data {
		keys = append(keys, k)
	}

	return keys
}

func (self *Params) Encode() string {
	var paramBody string

	if len(self.data) > 0 {
		var buf bytes.Buffer
		for k, v := range self.data {
			for _, vv := range v {
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(vv))
				buf.WriteByte('&')
			}
		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}

	return paramBody
}

func (self *Params) Get(key interface{}) (interface{}, bool) {

	for k, _ := range self.data {
		ks := key.(string)
		if k == ks {
			return self.data[k], true
		}
	}

	return nil, false
}
