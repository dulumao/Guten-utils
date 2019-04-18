package orderedmap

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func (m *OrderedMap) ComputeHmacSha256(secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(m.Encode()))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

func (m *OrderedMap) ComputeHmacSha256ToBase64String(secret string) string {
	return base64.StdEncoding.EncodeToString([]byte(m.ComputeHmacSha256(secret)))
}
