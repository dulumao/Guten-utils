package custom

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"errors"
)

func DecryptString(cryptoText string, keyString string) (plainTextString string) {
	newKeyString, err := hashTo32Bytes(keyString)
	cipherText, _ := base64.URLEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher([]byte(newKeyString))

	if err != nil {
		panic(err)
	}

	if len(cipherText) < aes.BlockSize {
		panic("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}

func EncryptString(plainText string, keyString string) (cipherTextString string) {
	newKeyString, err := hashTo32Bytes(keyString)

	if err != nil {
		panic(err)
	}

	key := []byte(newKeyString)
	value := []byte(plainText)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(value))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], value)

	return base64.URLEncoding.EncodeToString(cipherText)
}

func hashTo32Bytes(input string) (output string, err error) {

	if len(input) == 0 {
		return "", errors.New("No input supplied")
	}

	hasher := sha256.New()
	hasher.Write([]byte(input))
	stringToSHA256 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return stringToSHA256[:32], nil
}
