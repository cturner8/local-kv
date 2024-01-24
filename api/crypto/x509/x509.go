package x509

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func EncodePrivateKey(privateKey *rsa.PrivateKey) []byte {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})
	return pemBlock
}

func EncodePublicKey(publicKey *rsa.PublicKey) []byte {
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	pemBlock := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes})
	return pemBlock
}

func DecodePrivateKey(pemBlock []byte) *rsa.PrivateKey {
	privateKeyBlock, _ := pem.Decode(pemBlock)
	privateKey, _ := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	return privateKey
}

func DecodePublicKey(pemBlock []byte) *rsa.PublicKey {
	publicKeyBlock, _ := pem.Decode(pemBlock)
	publicKey, _ := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	return publicKey
}
