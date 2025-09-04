package cipher

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"
)

type Cipher struct {
	key []byte
}

// Returns new cipher
func NewCipher(key []byte) *Cipher {
	return &Cipher{key: key}
}

// Encrypts data
func (c *Cipher) Encrypt(data []byte) ([]byte, error) {
	if len(c.key) != 32 {
		return nil, errors.New("the key must be 32 bytes")
	}

	aead, err := chacha20poly1305.New(c.key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	rand.Read(nonce)

	ciphertext := aead.Seal(nil, nonce, data, nil)
	return append(nonce, ciphertext...), nil
}

// Decrypts data
func (c *Cipher) Decrypt(data []byte) ([]byte, error) {
	if len(c.key) != 32 || len(data) < 12 {
		return nil, errors.New("incorrect input data")
	}

	aead, err := chacha20poly1305.New(c.key)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:12], data[12:]
	return aead.Open(nil, nonce, ciphertext, nil)
}
