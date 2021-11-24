package gomen

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// EncryptString ...
func EncryptString(cryptKey string, text string) (string, error) {
	encryptText, err := Encrypt(cryptKey, []byte(text))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", encryptText), nil
}

// DecryptString ...
func DecryptString(cryptKey string, text string) (string, error) {
	b, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	decryptText, err := Decrypt(cryptKey, b)
	if err != nil {
		return "", err
	}
	return string(decryptText), nil
}

// Block Encrypt / Decrypt

// Encrypt ...
func Encrypt(cryptKey string, plainText []byte) ([]byte, error) {
	key := []byte(cryptKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	// fmt.Printf("Encrypt text: %x \n", cipherText)
	return cipherText, nil
}

// Decrypt ...
func Decrypt(cryptKey string, cipherText []byte) ([]byte, error) {
	key := []byte(cryptKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decryptedText := make([]byte, len(cipherText[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(block, cipherText[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedText, cipherText[aes.BlockSize:])
	// fmt.Printf("Decrypted text: %s\n", string(decryptedText))
	return decryptedText, nil
}
