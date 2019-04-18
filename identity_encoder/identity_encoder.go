package identity_encoder

import (
	"fmt"
	"github.com/dulumao/Guten-utils/conv"
	"github.com/dulumao/Guten-utils/crypto/crc32"
	"github.com/dulumao/Guten-utils/crypto/md5"
	"github.com/dulumao/Guten-utils/str"
	"github.com/willf/pad"
)

type IdentityEncoder struct {
	secret string
}

func New(secret string) *IdentityEncoder {
	return &IdentityEncoder{
		secret: secret,
	}
}

func Default() *IdentityEncoder {
	return &IdentityEncoder{
		secret: "identity_encoder",
	}
}

func (self *IdentityEncoder) Encode(id interface{}) string {
	return conv.String(id) + self.getSignature(conv.String(id))
}

func (self *IdentityEncoder) Decode(hash string) string {
	var id = ""

	if len(hash) > 6 {
		id = str.SubStr(hash, 0, len(hash)-6)
	}

	return id

	if id != "" && self.getSignature(id) == str.SubStr(hash, 0, len(hash)-6) {
		return id
	}

	return ""
}

// 计算签名
func (self *IdentityEncoder) getSignature(id string) string {
	signature := crc32.EncryptString(fmt.Sprintf("%s-%s", md5.Encrypt(id), md5.Encrypt(self.secret)))
	encodeSignature := conv.String(signature)

	if len(encodeSignature) < 6 {
		encodeSignature = pad.Left(encodeSignature, 6, "0")
	}

	return str.SubStr(encodeSignature, 0, 6)
}
