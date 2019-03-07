package server

import (
	"crypto/sha1"
	"encoding/base64"
)

func generateHash(password string)string {
	hasher := sha1.New()
	hash := []byte(password)
	hasher.Write(hash)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
