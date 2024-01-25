package x509

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestEncodePrivateKey(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	encoded := EncodePrivateKey(privateKey)

	if encoded == nil {
		t.Error("Expected encoded private key to be generated")
	}

	if len(encoded) == 0 {
		t.Error("Expected encoded private key to not be empty")
	}
}

func TestEncodePublicKey(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	encoded := EncodePublicKey(&privateKey.PublicKey)

	if encoded == nil {
		t.Error("Expected encoded public key to be generated")
	}

	if len(encoded) == 0 {
		t.Error("Expected encoded public key to not be empty")
	}
}

func TestDecodePrivateKey(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	encoded := EncodePrivateKey(privateKey)

	decoded := DecodePrivateKey(encoded)

	if decoded == nil {
		t.Error("Expected decoded private key to be generated")
	}

	if !decoded.Equal(privateKey) {
		t.Error("Expected matching private key to be generated")
	}
}

func TestDecodePublicKey(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	encoded := EncodePublicKey(&privateKey.PublicKey)

	decoded := DecodePublicKey(encoded)

	if decoded == nil {
		t.Error("Expected decoded public key to be generated")
	}

	if !decoded.Equal(&privateKey.PublicKey) {
		t.Error("Expected matching public key to be generated")
	}
}
