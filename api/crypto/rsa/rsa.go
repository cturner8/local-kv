package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}

func GenerateHash() hash.Hash {
	return sha256.New()
}

func Encrypt(publicKey *rsa.PublicKey, plaintext []byte, label []byte) string {
	hash := GenerateHash()
	// Encrypt the plaintext
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, plaintext, label)
	if err != nil {
		panic(err)
	}
	// Base64-encode the ciphertext
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(privateKey *rsa.PrivateKey, ciphertextBase64 string, label []byte) []byte {
	// Decode the base64-encoded data
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		panic(err)
	}

	hash := GenerateHash()
	// Decrypt the ciphertext
	plaintext, err := rsa.DecryptOAEP(hash, nil, privateKey, ciphertext, label)
	if err != nil {
		panic(err)
	}

	return plaintext
}
