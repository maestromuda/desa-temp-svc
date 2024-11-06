package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptStringToSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
