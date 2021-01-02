package suprlib

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
)

func MD5(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSHA256(msg []byte, secretKey []byte) ([]byte, error) {
	h := hmac.New(sha256.New, secretKey)
	_, err := h.Write(msg)
	if err != nil {
		return nil, err
	}
	hmacSha256 := h.Sum(nil)
	return hmacSha256, nil
}

func GZipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
