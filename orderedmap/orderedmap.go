package orderedmap

import (
	"bytes"
	"encoding/json"
	"sort"
	"sync"
)

func New() *OrderedMap {
	return &OrderedMap{
		store: make(map[interface{}]interface{}),
		keys:  make([]interface{}, 0, 0),
	}
}

type OrderedMap struct {
	sync.Mutex

	store map[interface{}]interface{}
	keys  []interface{}
}

func (m *OrderedMap) Set(key, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

func (m *OrderedMap) Get(key interface{}) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

func (m *OrderedMap) GetString(key interface{}) (string, bool) {
	value, found := m.Get(key)

	if found {
		if v, ok := value.(string); ok {
			return v, true
		}
	}

	return "", false
}

func (m *OrderedMap) MustGet(key interface{}) interface{} {
	value, found := m.Get(key)

	if found {
		return value
	}

	return nil
}

func (m *OrderedMap) MustGetString(key interface{}) string {
	value, found := m.GetString(key)

	if found {
		return value
	}

	return ""
}

// Remove remove a key-value pair from the OrderedMap
func (m *OrderedMap) Delete(key interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		return
	}

	delete(m.store, key)

	for i, _ := range m.keys {
		if m.keys[i] == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *OrderedMap) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

func (m *OrderedMap) Keys() []interface{} {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

func (m *OrderedMap) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

func (m *OrderedMap) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}

func (m *OrderedMap) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	var idx = 0

	for _, key := range m.Keys() {
		jsonKey, err := json.Marshal(key)

		if err != nil {
			return nil, err
		}

		jsonValue, err := json.Marshal(m.store[key])

		if err != nil {
			return nil, err
		}

		buffer.Write(jsonKey)
		buffer.WriteByte(58)
		buffer.Write(jsonValue)

		idx++

		if idx != m.Size() {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString("}")

	return buffer.Bytes(), nil
}

func (m *OrderedMap) String() string {
	json, _ := m.MarshalJSON()

	return string(json)
}

func (m *OrderedMap) SortKeyByString() {
	m.Lock()
	defer m.Unlock()

	var alphabetKeys []string

	for _, i := range m.keys {
		alphabetKeys = append(alphabetKeys, i.(string))
	}

	sort.Strings(alphabetKeys)

	var replaceKeys []interface{}

	for _, i := range alphabetKeys {
		replaceKeys = append(replaceKeys, i)
	}

	m.keys = replaceKeys
}

/*
m.SortKeys(sort.Strings)
*/
func (m *OrderedMap) SortKeys(sortFunc interface{}) {
	m.Lock()
	defer m.Unlock()

	var replaceKeys []interface{}

	switch sortFunc.(type) {
	case func(keys []string):
		var originKeys []string

		for _, i := range m.keys {
			originKeys = append(originKeys, i.(string))
		}

		sortFunc.(func(keys []string))(originKeys)

		for _, i := range originKeys {
			replaceKeys = append(replaceKeys, i)
		}
	case func(keys []int):
		var originKeys []int

		for _, i := range m.keys {
			originKeys = append(originKeys, i.(int))
		}

		sortFunc.(func(keys []int))(originKeys)

		for _, i := range originKeys {
			replaceKeys = append(replaceKeys, i)
		}
	case func(keys []float64):
		var originKeys []float64

		for _, i := range m.keys {
			originKeys = append(originKeys, i.(float64))
		}

		sortFunc.(func(keys []float64))(originKeys)

		for _, i := range originKeys {
			replaceKeys = append(replaceKeys, i)
		}
	}

	m.keys = replaceKeys
}
