package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

const (
	passwordEncryptSeed = "(beyond)@#$"
	mobileAesKey        = "5A2E746B08D846502F37A6E2D85D583B"
)

func EncPassword(password string) string {
	return Md5Sum([]byte(strings.TrimSpace(password + passwordEncryptSeed)))
}

func Md5Sum(data []byte) string {
	return hex.EncodeToString(byte16ToBytes(md5.Sum(data)))
}
func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	return tmp[16:]
}
