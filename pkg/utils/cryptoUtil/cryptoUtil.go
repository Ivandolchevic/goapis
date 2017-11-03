package cryptoUtil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

var key = []byte("4f23f69eb995405db2530c3902b1e653")
var noncestr = "37b8e8a308c354048d245f6d"

// Encrypt crypts a value using AES 256
func Encrypt(value string) (string, error) {
	plaintext := []byte(value)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return string(ciphertext), nil
}

// Decrypt decrypts a value encrypted using AES 256
func Decrypt(value string) (string, error) {
	ciphertext, _ := hex.DecodeString(value)

	nonce, _ := hex.DecodeString(noncestr)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
