package orderedmap

import (
	"github.com/dulumao/Guten-utils/conv"
	"strings"
)

func (m *OrderedMap) JoinKV() string {
	if m == nil {
		return ""
	}

	var buf strings.Builder

	for _, k := range m.Keys() {
		v, _ := m.Get(k)

		// buf.WriteString(url.QueryEscape(conv.String(k)))
		// buf.WriteString(url.QueryEscape(conv.String(v)))
		buf.WriteString(conv.String(k))
		buf.WriteString(conv.String(v))
	}

	return buf.String()
}
