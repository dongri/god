package barikata

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"testing"
)

func TestSha1Sum(t *testing.T) {
	s1 := Sha1Sum("hello")
	bytes := sha1.Sum([]byte("hello"))
	s2 := hex.EncodeToString(bytes[:])
	if s1 != s2 {
		t.Fatal("Error!")
	}
}

func TestSha256Sum224(t *testing.T) {
	s1 := Sha256Sum224("hello")
	bytes := sha256.Sum224([]byte("hello"))
	s2 := hex.EncodeToString(bytes[:])
	if s1 != s2 {
		t.Fatal("Error!")
	}
}

func TestSha256Sum256(t *testing.T) {
	s1 := Sha256Sum256("hello")
	bytes := sha256.Sum256([]byte("hello"))
	s2 := hex.EncodeToString(bytes[:])
	if s1 != s2 {
		t.Fatal("Error!")
	}
}

func TestSha512Sum384(t *testing.T) {
	s1 := Sha512Sum384("hello")
	bytes := sha512.Sum384([]byte("hello"))
	s2 := hex.EncodeToString(bytes[:])
	if s1 != s2 {
		t.Fatal("Error!")
	}
}

func TestSha512Sum512(t *testing.T) {
	s1 := Sha512Sum512("hello")
	bytes := sha512.Sum512([]byte("hello"))
	s2 := hex.EncodeToString(bytes[:])
	if s1 != s2 {
		t.Fatal("Error!")
	}
}
