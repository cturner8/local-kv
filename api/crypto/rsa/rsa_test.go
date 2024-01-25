package rsa

import (
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	privateKey, publicKey := GenerateKeyPair(2048)

	if privateKey == nil {
		t.Error("Expected private key to be generated")
	}

	if publicKey == nil {
		t.Error("Expected public key to be generated")
	}
}

func TestGenerateHash(t *testing.T) {
	hash := GenerateHash()
	if hash == nil {
		t.Error("Expected hash to be generated")
	}
}

func TestEncrypt(t *testing.T) {
	_, publicKey := GenerateKeyPair(2048)

	plaintext := "lorem ipsum dolor sit amet"
	ciphertext := Encrypt(publicKey, []byte(plaintext), nil)
	if len(ciphertext) == 0 {
		t.Errorf("Expected non-empty ciphertext")
	}
}

func TestEncryptWithLabel(t *testing.T) {
	_, publicKey := GenerateKeyPair(2048)

	plaintext := "lorem ipsum dolor sit amet"
	ciphertext := Encrypt(publicKey, []byte(plaintext), []byte("john doe"))
	if len(ciphertext) == 0 {
		t.Errorf("Expected non-empty ciphertext")
	}
}

func TestDecrypt(t *testing.T) {
	privateKey, publicKey := GenerateKeyPair(2048)

	plaintext := "lorem ipsum dolor sit amet"
	ciphertext := Encrypt(publicKey, []byte(plaintext), nil)
	decryptedCiphertext := Decrypt(privateKey, ciphertext, nil)

	if string(decryptedCiphertext) != plaintext {
		t.Error("Expected matching plaintext and decrypted ciphertext")
	}
}

func TestDecryptWithLabel(t *testing.T) {
	privateKey, publicKey := GenerateKeyPair(2048)

	plaintext := "lorem ipsum dolor sit amet"
	label := []byte("john doe")
	ciphertext := Encrypt(publicKey, []byte(plaintext), label)
	decryptedCiphertext := Decrypt(privateKey, ciphertext, label)

	if string(decryptedCiphertext) != plaintext {
		t.Error("Expected matching plaintext and decrypted ciphertext")
	}
}

func TestDecryptWithWrongKey(t *testing.T) {
	defer func() {
		err := recover()

		if err == nil {
			t.Error("Expected decrypt process to error")
		}
	}()

	_, publicKey := GenerateKeyPair(2048)

	plaintext := "lorem ipsum dolor sit amet"
	label := []byte("john doe")
	ciphertext := Encrypt(publicKey, []byte(plaintext), label)

	privateKey, _ := GenerateKeyPair(2048)
	Decrypt(privateKey, ciphertext, label)

	t.Error("Expected failure")
}

func TestDecryptWithWrongLabel(t *testing.T) {
	defer func() {
		err := recover()

		if err == nil {
			t.Error("Expected decrypt process to error")
		}
	}()

	privateKey, publicKey := GenerateKeyPair(2048)

	plaintext := "lorem ipsum dolor sit amet"
	label := []byte("john doe")
	ciphertext := Encrypt(publicKey, []byte(plaintext), label)

	Decrypt(privateKey, ciphertext, nil)

	t.Error("Expected failure")
}
