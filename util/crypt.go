package util

import (
	"crypto/md5"
	"crypto/rc4"
	"encoding/hex"
	"net/url"
)

func EncryptMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptMd5Byte(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func UrlEncode(str string) string {

	return url.QueryEscape(str)

}

func UrlDecode(str string) (string, error) {

	return url.QueryUnescape(str)

}

func DecryptRC4(str string, key string) ([]byte, error) {

	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(str))
	cipher.XORKeyStream(dst, []byte(str))

	return dst, nil
}

func DecryptRC4Byte(str []byte, key string) ([]byte, error) {

	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(str))
	cipher.XORKeyStream(dst, str)

	return dst, nil
}
