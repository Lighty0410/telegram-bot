package server

import (
	"crypto/sha1"
	"encoding/base64"
	"log"
)

// generateHash generates hash for password.
func generateHash(password string) string {
	hasher := sha1.New()
	hash := []byte(password)
	_, err := hasher.Write(hash)
	if err != nil {
		log.Println(err)
	}
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
