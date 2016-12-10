package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Sha1Sum ...
func Sha1Sum(s string) string {
	data := []byte(s)
	bytes := sha1.Sum(data)
	return hex.EncodeToString(bytes[:])
}

// Sha256Sum224 ...
func Sha256Sum224(s string) string {
	data := []byte(s)
	bytes := sha256.Sum224(data)
	return hex.EncodeToString(bytes[:])
}

// Sha256Sum256 ...
func Sha256Sum256(s string) string {
	data := []byte(s)
	bytes := sha256.Sum256(data)
	return hex.EncodeToString(bytes[:])
}

// Sha512Sum384 ...
func Sha512Sum384(s string) string {
	data := []byte(s)
	bytes := sha512.Sum384(data)
	return hex.EncodeToString(bytes[:])
}

// Sha512Sum512 ...
func Sha512Sum512(s string) string {
	data := []byte(s)
	bytes := sha512.Sum512(data)
	return hex.EncodeToString(bytes[:])
}
