package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

// GetMd5String 生成32位MD5字符串
func GetMd5String(s string) string {
	newmd5 := md5.New()
	newmd5.Write([]byte(s))
	return hex.EncodeToString(newmd5.Sum(nil))
}

// GetUniqueID 生成UID唯一标识
func GetUniqueID() (string, error) {
	newbyte := make([]byte, 48)

	_, err := io.ReadFull(rand.Reader, newbyte)
	if err != nil {
		return "", err
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(newbyte)), nil
}
