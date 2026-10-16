package security

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

// GenerateSalt generates a salt in string
func GenerateSalt() (string, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%x",
		b,
	), nil
}

// ToHash converts password to has
func ToHash(
	password string,
	salt string,
) string {
	hashInstance := sha1.New()
	hashInstance.Write([]byte(password + salt))
	result := hex.EncodeToString(
		hashInstance.Sum(nil),
	)
	return result
}

// ComparePassword compares the hash value and password get
func ComparePassword(
	password string,
	salt string,
	storedHash string,
) bool {
	hashValue := ToHash(password, salt)
	return hashValue == storedHash
}
