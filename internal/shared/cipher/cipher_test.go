package cipher

import (
	"reflect"
	"testing"
)

func TestNewCipher(t *testing.T) {
	t.Parallel() // Run in parallel with other tests

	tests := []struct {
		name string
		key  []byte
	}{
		{
			name: "non-empty key",
			key:  []byte("my-secret-key-32bytes-1234567890"),
		},
		{
			name: "empty slice (len=0)",
			key:  []byte{},
		},
		{
			name: "nil slice",
			key:  nil,
		},
		{
			name: "single byte key",
			key:  []byte{42},
		},
		{
			name: "multi-byte key",
			key:  []byte{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipher := NewCipher(tt.key)

			if cipher == nil {
				t.Fatalf("NewCipher() returned nil for key %v", tt.key)
			}

			if !reflect.DeepEqual(cipher.key, tt.key) {
				t.Errorf("cipher.key = %v, want %v", cipher.key, tt.key)
			}

			if tt.key != nil && len(tt.key) > 0 {
				modifiedKey := make([]byte, len(tt.key))
				copy(modifiedKey, tt.key)
				modifiedKey[0] = ^modifiedKey[0]

				if reflect.DeepEqual(cipher.key, modifiedKey) {
					t.Error("cipher.key appears to alias input; expected independent copy")
				}
			}
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i)
	}

	plaintext := []byte("Hello")

	encrypted, err := NewCipher(key).Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := NewCipher(key).Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if !reflect.DeepEqual(decrypted, plaintext) {
		t.Errorf("The decrypted data does not match. Received: %v, expected: %v", decrypted, plaintext)
	}
}

func TestInvalidKey(t *testing.T) {
	_, err := NewCipher([]byte("short")).Encrypt([]byte("test"))
	if err == nil {
		t.Error("Encryption should fail with a short key")
	}
}
