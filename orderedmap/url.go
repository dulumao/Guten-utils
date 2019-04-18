package orderedmap

import (
	"github.com/dulumao/Guten-utils/conv"
	"net/url"
	"strings"
)

func (m *OrderedMap) Encode() string {
	if m == nil {
		return ""
	}

	var buf strings.Builder

	for _, k := range m.Keys() {
		v, _ := m.Get(k)

		keyEscaped := url.QueryEscape(conv.String(k))

		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(conv.String(v)))
	}

	return buf.String()
}
