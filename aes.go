package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

func Encrypt(data []byte, key []byte) ([]byte, error) {
	// Keys must be 32 bytes long, so we use SHA-256 to hash the key
	hashedKey := sha256.Sum256(key)
	c, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	// Keys must be 32 bytes long, so we use SHA-256 to hash the key
	hashedKey := sha256.Sum256(key)
	c, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
