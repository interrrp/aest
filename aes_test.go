package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	data := []byte("hello")
	key := []byte("correct horse battery staple")
	enc, err := Encrypt(data, key)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, data, enc)
}

func TestDecrypt(t *testing.T) {
	data := []byte("hello")
	key := []byte("correct horse battery staple")
	enc, err := Encrypt(data, key)
	if err != nil {
		t.Fatal(err)
	}
	dec, err := Decrypt(enc, key)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, data, dec)
}
