package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

func GenerateDataKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}

func GenerateNonce() []byte {
	nonce := make([]byte, 12)
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err)
	}
	return nonce
}

func GenerateCipher(key []byte) cipher.AEAD {
	// Create a cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Create a GCM cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	return aesgcm
}

func Encrypt(key []byte, plaintext []byte, additionalData *[]byte) string {
	aesgcm := GenerateCipher(key)

	// Encrypt the plaintext
	nonce := GenerateNonce()
	var ciphertext []byte
	if additionalData != nil {
		ciphertext = aesgcm.Seal(nil, nonce, plaintext, *additionalData)
	} else {
		ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	}

	// Combine nonce and ciphertext
	ciphertextWithNonce := append(nonce, ciphertext...)

	// Base64-encode the combined data
	return base64.StdEncoding.EncodeToString(ciphertextWithNonce)
}

func Decrypt(key []byte, ciphertextBase64 string, additionalData *[]byte) []byte {
	var nonce, ciphertext []byte
	aesgcm := GenerateCipher(key)

	// Decode the base64-encoded data
	ciphertextWithNonce, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		panic(err)
	}

	// Separate nonce and ciphertext
	nonce = ciphertextWithNonce[:12]
	ciphertext = ciphertextWithNonce[12:]

	// Decrypt the ciphertext
	var plaintext []byte
	if additionalData != nil {
		plaintext, err = aesgcm.Open(nil, nonce, ciphertext, *additionalData)
	} else {
		plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	}

	if err != nil {
		panic(err)
	}

	return plaintext
}
