package aes

import (
	"bytes"
	"testing"
)

func TestGenerateDataKey(t *testing.T) {
	key := GenerateDataKey()
	if len(key) != 32 {
		t.Errorf("Expected key length of 32, got %d", len(key))
	}
}

func TestGenerateNonce(t *testing.T) {
	nonce := GenerateNonce()
	if len(nonce) != 12 {
		t.Errorf("Expected nonce length of 12, got %d", len(nonce))
	}
}

func TestGenerateUniqueNonce(t *testing.T) {
	nonce1 := GenerateNonce()
	nonce2 := GenerateNonce()
	if bytes.Equal(nonce1, nonce2) {
		t.Errorf("Expected nonces to be unique")
	}
}

func TestEncypt(t *testing.T) {
	key := GenerateDataKey()
	plaintext := "lorem ipsum dolor sit amet"
	ciphertext := Encrypt(key, []byte(plaintext), nil)
	if len(ciphertext) == 0 {
		t.Errorf("Expected ciphertext to not be empty")
	}
}

func TestDecrypt(t *testing.T) {
	key := GenerateDataKey()
	plaintext := "lorem ipsum dolor sit amet"
	ciphertext := Encrypt(key, []byte(plaintext), nil)
	decryptedPlaintext := Decrypt(key, ciphertext, nil)
	if string(decryptedPlaintext) != plaintext {
		t.Errorf("Expected decrypted plaintext to be %s, got %s", plaintext, string(decryptedPlaintext))
	}
}

func TestEncyptWithAad(t *testing.T) {
	key := GenerateDataKey()
	plaintext := "lorem ipsum dolor sit amet"
	additionalData := []byte("john doe")
	ciphertext := Encrypt(key, []byte(plaintext), &additionalData)
	if len(ciphertext) == 0 {
		t.Errorf("Expected ciphertext to not be empty")
	}
}

func TestDecryptWithAad(t *testing.T) {
	key := GenerateDataKey()
	plaintext := "lorem ipsum dolor sit amet"
	additionalData := []byte("john doe")
	ciphertext := Encrypt(key, []byte(plaintext), &additionalData)
	decryptedPlaintext := Decrypt(key, ciphertext, &additionalData)
	if string(decryptedPlaintext) != plaintext {
		t.Errorf("Expected decrypted plaintext to be %s, got %s", plaintext, string(decryptedPlaintext))
	}
}

func TestDecryptWithWrongKey(t *testing.T) {
	defer func() {
		err := recover()

		if err == nil {
			t.Error("Expected decrypt process to error")
		}
	}()

	key := GenerateDataKey()

	plaintext := "lorem ipsum dolor sit amet"
	ciphertext := Encrypt(key, []byte(plaintext), nil)

	key = GenerateDataKey()
	Decrypt(key, ciphertext, nil)

	t.Error("Expected failure")
}

func TestDecryptWithWrongAad(t *testing.T) {
	defer func() {
		err := recover()

		if err == nil {
			t.Error("Expected decrypt process to error")
		}
	}()

	key := GenerateDataKey()

	plaintext := "lorem ipsum dolor sit amet"
	additionalData := []byte("john doe")
	ciphertext := Encrypt(key, []byte(plaintext), &additionalData)

	Decrypt(key, ciphertext, nil)

	t.Error("Expected failure")
}
